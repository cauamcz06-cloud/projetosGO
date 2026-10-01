package Main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	password := geratePassword(5)
	fmt.Println(password)
}
func geratePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%¨&*()+?><:{}[]"
	allChars := lowerCase + upperCase + numbers + special
	mandatory := []byte{
		upperCase[rand.IntN(len(upperCase))],
		numbers[rand.IntN(len(numbers))],
	}

	password := make([]byte, length-len(mandatory))
	for i := range password {
		password[i] = allChars[rand.IntN(len(allChars))]
	}
	password = append(password, mandatory...)
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]x
	})

	return string(password)
}
