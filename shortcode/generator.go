package shortcode

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLen = 8
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Generate() string {
	code := make([]byte, codeLen)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func GenerateUnique(checkExixts func(string) bool) string {
	for {
		code := Generate()
		if !checkExixts(code) {
			return code
		}
	}
}