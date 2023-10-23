package utility

import (
	"math/rand"
	"strconv"
)

func GenerateDigitToken(digits int) string {
	id := ""
	for i := 0; i < digits; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}
	return id
}
