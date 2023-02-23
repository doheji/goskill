package utils

import (
	"math/rand"
	"strings"
)

var alphabets = []byte("abcdefghijklmnopqrstuvwxyz")

func RandomNumber(min, max int) int64 {

	return rand.Int63n(int64(max-min+1)) + int64(min)
}

func RandomName(n int) string {
	builder := strings.Builder{}
	k := len(alphabets)
	for i := 0; i < n; i++ {
		builder.WriteByte(alphabets[rand.Intn(k)])
	}
	return builder.String()
}
