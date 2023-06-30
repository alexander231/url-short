package base62

import "strings"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Encode(n uint64) string {
	length := uint64(len(alphabet))
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)

	for n > 0 {
		remainder := n % length
		encodedBuilder.WriteByte(alphabet[remainder])
		n = n / length
	}

	return encodedBuilder.String()
}
