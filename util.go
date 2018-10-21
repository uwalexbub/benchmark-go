package main

import (
	"math/rand"
	"time"
)

type delegate func()

func initRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

// This function is sources from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func getUniqueValue() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 6
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func measureAction(action delegate) int64 {
	start := time.Now()
	action()
	elapsedNanos := time.Now().Sub(start).Nanoseconds()

	return elapsedNanos
}

func avg(sum int64, n int) int64 {
	return sum / int64(n)
}

func dummyAction(i int, s string) {

}
