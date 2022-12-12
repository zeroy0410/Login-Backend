package utility

import (
	// "fmt"
	"math/rand"
	"strconv"
	"time"
)

var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func AtoU(str string) (uint, error) {
	i, err := strconv.Atoi(str)
	return uint(i), err
}

func Initialize() error {
	rand.Seed(time.Now().Unix())
	return nil
}

func GetRandString() (string, error) {
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		code[i] = longLetters[rand.Int()%62]
	}
	return string(code), nil
}
