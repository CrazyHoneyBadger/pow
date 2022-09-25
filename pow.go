package pow

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"golang.org/x/crypto/sha3"
)

const (
	VERSION = "0.1.0"
)

var (
	ErrVersion = fmt.Errorf("versions pow library on the server and client are different")
	ErrHash    = fmt.Errorf("hash is not valid")
)

func generate512Hash(message string) string {
	hash := sha512.New()
	hash.Write([]byte(message))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func generate3Hash(message string) string {
	hash := sha3.New512()
	hash.Write([]byte(message))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func hashValidate(hash string, complexity int) bool {
	preifx := strings.Repeat("0", complexity)
	return strings.HasPrefix(hash, preifx)
}
