package utils

import "math/rand"

func RandStringBytes(n int, charRange string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charRange[rand.Intn(len(charRange))]
	}
	return string(b)
}
