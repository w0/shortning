package base62

import (
	"math"
	"strings"
)

const symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Encode(number int) string {
	encoded := ""

	length := len(symbols)
	for number > 0 {
		rem := number % length
		encoded = string(symbols[rem]) + encoded
		number /= length
	}

	return encoded
}

func Decode(s string) int {
	trimed := strings.TrimPrefix(s, "/")

	decoded := 0
	pow := len(trimed) - 1
	for i := 0; i < len(trimed); i++ {
		idx := strings.Index(symbols, string(trimed[i]))
		decoded += idx * int(math.Pow(float64(62), float64(pow)))
		pow--
	}

	return decoded
}
