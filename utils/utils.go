package utils

import (
	"math/rand"
	"time"
)

func RandomArray(size uint) []byte {
	arr := make([]byte, size)

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(arr); i++ {
		c := rand.Intn(int('{'))
		if c >= 'A' && c <= 'z' {
			arr[i] = byte(c)
		} else {
			i--
		}
	}

	return arr
}
