package utils

import (
	"math/rand"
	"time"
)

// RandomSleep sleep random time
func RandomSleep() {
	randValue := 7 + rand.Int()%4
	time.Sleep(time.Duration(randValue) * time.Second)
}
