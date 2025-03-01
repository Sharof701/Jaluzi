package helper

import (
	"regexp"
)

// IsValidPhone ...
func IsValidPhone(phone string) bool {
	r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phone)
}

// IsValidUUID ...
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// IsValidPassword ...
func ValidatePassword(password string) bool {
	r := regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#~$%^&*()_+={}\[\]:;"'<>,.?\/\\|` + "`" + `-]).{8,}$`)
	return r.MatchString(password)
}
