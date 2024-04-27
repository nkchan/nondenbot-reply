package main
import (
	"math/rand"
	"time"
)

func RandomChoice(data []string) string {
	rand.Seed(time.Now().UnixNano())
	z := rand.Intn(len(data))
	return data[z]
}
