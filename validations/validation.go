package validations

import (
	"errors"
	"main/structs"
	"regexp"
)

func ValidateUser(user structs.User) (bool, error) {
	if !mobileValidation(user.MobileNumber) {
		return false, errors.New("invalid mobile")
	}

	if !nameValidation(user.FirstName) {
		return false, errors.New("invalid first name")
	}

	if !nameValidation(user.LastName) {
		return false, errors.New("invalid last name")
	}

	return true, nil

}

func mobileValidation(mobile string) bool {
	// regular expression pattern for mobile
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(mobile)
}

func nameValidation(name string) bool {
	re := regexp.MustCompile(`^[A-Za-z\s'-]+$`)
	return re.MatchString(name)
}
