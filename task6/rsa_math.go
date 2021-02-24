package main

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math/big"
)

func GCD(a int64, b int64) int64 {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func EGCD(a int64, b int64) (int64, int64, int64) {
	var x, xx, y, yy int64 = 1, 0, 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a % b
		x, xx = xx, x - xx * q
		y, yy = yy, y - yy * q
	}
	return x, y, a
}

func InverseByModulo(a int64, m int64) (int64, error)  {
	x, _, d := EGCD(a, m)
	if d != 1 {
		return 0, errors.New("NO INVERSE")
	}
	return (x % m + m) % m, nil
}

// only for n <= 62 :(
func PrimalityFermat(n int64) (bool, error) {
	if n <= 1 || n >= 63 {
		return false, errors.New("BAD N")
	}
	if n == 1 {
		return false, nil
	}
	if n == 2 {
		return true, nil
	}
	p := 1 << n
	return p % 2 == 1, nil
}

// for n <= 9_223_372_036_854_775_807 == 2 ** 63
func Primality(n int64) (bool, int64, error) {
	if n < 0 {
		return false, 0, errors.New("BAD N")
	}
	if n <= 1 {
		return false, 0, nil
	}
	if n == 2 {
		return true, 0, nil
	}
	if n % 2 == 0 {
		return false, 2, nil
	}
	for i := int64(3); i * i <= n; i += 2 {
		if n % i == 0 {
			return false, i, nil
		}
	}
	return true, 1, nil
}

func GenRandomBytes(l int) ([]byte, error){
	b := make([]byte, l)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenRandomInteger(bytesLen int) int64 {
	b, err := GenRandomBytes(bytesLen)
	for err != nil {
		b, err = GenRandomBytes(bytesLen)
	}
	b = append(make([]byte, 8 - bytesLen), b...)
	return int64(binary.BigEndian.Uint64(b))
}

func GenPrimeNumber(bytesLen int) int64 {
	for {
		n := GenRandomInteger(bytesLen)
		isPrime, _, _ := Primality(n)
		if isPrime {
			return n
		}
	}
}

func BigIntsToString(encoded []big.Int) (bigs []string) {
	bigs = []string{}
	for _, c := range encoded {
		bigs = append(bigs, c.String())
	}
	return
}

func GetRunes(s string) (runes []int32) {
	runes = []int32{}
	for _, r := range s {
		runes = append(runes, r)
	}
	return
}
