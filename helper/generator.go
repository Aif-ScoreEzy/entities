package helper

import (
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	lowerCharSet = "abcdedfghijklmnopqrst"
	upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet    = "0123456789"
	// allCharSet   = lowerCharSet + upperCharSet + numberSet
)

// Hash Use to hash password
func Hash(pwd string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashed)
}

// GenerateAPIKey ...
func GenerateAPIKey() string {
	length := 32
	number := 11
	upper_case := 11
	return GenerateRandomKey(length, number, upper_case)
}

// GenerateRandomKey Generate random key for APIKeys
func GenerateRandomKey(length, number, upper_case int) string {
	var password strings.Builder

	//Set numerice
	for i := 0; i <= number; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < upper_case; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := length - number - upper_case
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
