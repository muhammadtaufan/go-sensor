package utils

import (
	"math/rand"
	"time"
)

func GetRandomFloat() float32 {
	rand.Seed(time.Now().UnixNano())

	return rand.Float32() * 10
}

func GetRandomAlphabets() string {
	rand.Seed(time.Now().UnixNano())

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomChar := charset[rand.Intn(len(charset))]

	return string(randomChar)
}
