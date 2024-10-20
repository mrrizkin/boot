package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func NanoID() (string, error) {
	return gonanoid.New()
}

func CustomNanoID(char string, size int) (string, error) {
	return gonanoid.Generate(char, size)
}
