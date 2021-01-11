package generator

import (
	"crypto/rand"
	"math/big"
)

var (
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func ByteArray(number int64, pattern string, errorHandler func(*error)) (data *[]byte) {
	var (
		array = make([]byte, number, number)
		i1    int64
	)
	for i1 < number {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
		errorHandler(&err)
		if n != nil {
			array[i1] = pattern[n.Int64()]
		}
		i1++
	}
	data = &array
	return
}

func String(number int64, pattern string, errorHandler func(*error)) (data string) {
	var (
		array = make([]byte, number, number)
		i1    int64
	)
	for i1 < number {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
		errorHandler(&err)
		if n != nil {
			array[i1] = pattern[n.Int64()]
		}
		i1++
	}
	data = string(array)
	return
}
