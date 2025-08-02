package auth

// Change this to a hashed check or keychain later
const hardcodedPassword = "mySecret123"

func CheckPassword(input string) bool {
	return input == hardcodedPassword
}
