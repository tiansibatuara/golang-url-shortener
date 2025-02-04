package utils

import "github.com/teris-io/shortid"

func GenerateShortCode() (string, error) {
	generator := shortid.MustNew(1, shortid.DefaultABC, 2342)
	return generator.Generate()
}
