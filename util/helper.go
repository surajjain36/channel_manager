package util

import (
	"math/rand"
)

//Helper ...
type Helper struct{}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// GenerateRandomString : It generates random string
func (h *Helper) GenerateRandomString(length int) (randomString string) {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
