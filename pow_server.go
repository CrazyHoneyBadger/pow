package pow

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~=+%^*/()[]{}/!@#$?|" // Не добовлять символы не из ASCII
)

type POWServer struct {
	lengthUniqueKey int
	complexity      int
}

func NewPOWServer(lengthUniqueKey, complexity int) *POWServer {
	if lengthUniqueKey < 0 || complexity < 0 {
		lengthUniqueKey = 0
		complexity = 0
	}
	return &POWServer{
		lengthUniqueKey: lengthUniqueKey,
		complexity:      complexity,
	}
}
func (s POWServer) GenerateUniqKey() string {
	if s.lengthUniqueKey == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	var uniqKey strings.Builder
	for i := 0; i < s.lengthUniqueKey; i++ {
		uniqKey.WriteByte(letters[rand.Intn(len(letters))]) //Если использовать utf-8 то необходимо менять
	}
	return uniqKey.String()
}
func (s POWServer) ValidateMessage(version, message string) error {
	if s.lengthUniqueKey == 0 || s.complexity == 0 {
		return nil
	}
	if VERSION != version {
		return ErrVersion
	}
	hash := generate3Hash(message)
	if !hashValidate(hash, s.complexity) {
		return ErrHash
	}
	return nil
}
func (s POWServer) GetComplexity() int {
	return s.complexity
}
func (s POWServer) GetVersion() string {
	return VERSION
}
