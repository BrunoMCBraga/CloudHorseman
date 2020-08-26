package util

import (
	"math/rand"
	"strconv"
	"time"
)

func InitRandomGeneratorSeed() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomTimestampString() string {

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(3000000))

}
