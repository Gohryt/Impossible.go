package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	//Letters is standart pattern for string generation
	Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

//ByteArray function generates random string as byte array
func ByteArray(number int, pattern string) (data []byte, err error) {
	if number > 0 {
		data = make([]byte, number)
		for i := 0; i < number; i++ {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
			if err != nil {
				return data, err
			}
			if n != nil {
				data[i] = pattern[n.Int64()]
			}
		}
	}

	return
}

//String function generates random string
func String(number int, pattern string) (data string, err error) {
	if number > 0 {
		array := make([]byte, number)
		for i := 0; i < number; i++ {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pattern))))
			if err != nil {
				return data, err
			}
			if n != nil {
				array[i] = pattern[n.Int64()]
			}
		}
		data = string(array)
	}

	return
}
