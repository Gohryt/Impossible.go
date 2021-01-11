package generator

import (
	"crypto/rand"
	"math/big"
)

var (
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func ByteArray(number int, pattern string, errorHandler func(*error)) (data *[]byte) {
	array := make([]byte, number, number)
	for i := 0; i < number; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
		errorHandler(&err)
		if n != nil {
			array[i] = pattern[n.Int64()]
		}
	}
	data = &array
	return
}

func String(number int, pattern string, errorHandler func(*error)) (data string) {
	array := make([]byte, number, number)
	for i := 0; i < number; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
		errorHandler(&err)
		if n != nil {
			array[i] = pattern[n.Int64()]
		}
	}
	data = string(array)
	return
}
