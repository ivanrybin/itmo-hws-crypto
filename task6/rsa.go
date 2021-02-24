package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

const e = int64(3)

func genRSANumbers(primeBytesLen int) (int64, int64, int64, error) {
	p := GenPrimeNumber(primeBytesLen)
	q := GenPrimeNumber(primeBytesLen)
	n := p * q
	phi := (p - 1) * (q - 1)
	for p == q || n < math.MaxInt32 || GCD(e, phi) != 1 {
		p = GenPrimeNumber(primeBytesLen)
		q = GenPrimeNumber(primeBytesLen)
		n = p * q
		phi = (p - 1) * (q - 1)
	}
	d, err := InverseByModulo(e, phi)
	if err != nil {
		return p, q, 0, err
	}
	if (d * e) % phi != 1 {
		return p, q, d, errors.New("BAD INVERSE D")
	}
	return p, q, d, nil
}

func encodeRSA(text string, p int64, q int64, salt int64) []big.Int {
	bigE := big.NewInt(e)
	bigN := big.NewInt(p * q)
	bigSalt := big.NewInt(salt)

	encoded := []big.Int{}
	for _, char := range text {
		bigChar := big.NewInt(int64(char))

		saltedChar := big.NewInt(salt)
		saltedChar = saltedChar.Add(bigChar, bigSalt)

		encodedChar := big.NewInt(0)
		encodedChar = encodedChar.Exp(saltedChar, bigE, bigN)

		encoded = append(encoded, *encodedChar)
	}
	return encoded
}

func decodeRSA(encoded []big.Int, p int64, q int64, d int64, salt int64) string {
	bigD := big.NewInt(d)
	bigN := big.NewInt(p * q)
	bigSalt := big.NewInt(salt)

	decoded := []int32{}
	for _, encodedChar := range encoded {
		decodedSaltedChar := big.NewInt(0)
		decodedSaltedChar = decodedSaltedChar.Exp(&encodedChar, bigD, bigN)

		decodedChar := big.NewInt(0)
		decodedChar = decodedSaltedChar.Sub(decodedSaltedChar, bigSalt)

		decoded = append(decoded, int32(decodedChar.Int64()))
	}
	return string(decoded)
}

func main() {
	p, q, d, err := genRSANumbers(4)
	for err != nil {
		p, q, d, err = genRSANumbers(4)
	}
	salt := int64(1000000)

	text := "Всем привет! 你好! Hi! Guten Tag!"
	encoded := encodeRSA(text, p, q, salt)
	decoded := decodeRSA(encoded, p, q, d, salt)

	fmt.Printf("e=%d\np=%d\nq=%d\nN=%d\nd=%d\n(d * e (mod (p - 1) * (q - 1))=%d\n",
		e, p, q, p * q, d, (e * d) % ((p - 1) * (q - 1)),
	)
	fmt.Println("text: ", text)
	fmt.Printf("source  (%d elems): %v\n", len(GetRunes(text)), GetRunes(text))
	fmt.Printf("encoded (%d elems): %v\n", len(encoded), BigIntsToString(encoded))
	fmt.Println("decoded: ", decoded)
}
