package val

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/PlatosCodes/momsrecipes/util"
)

var isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
var isValidRecipeName = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`).MatchString

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lower case letters, digits or underscores")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	_, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("not a valid email address")
	}
	return nil
}

func ValidateRecipeName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidRecipeName(value) {
		return fmt.Errorf("must contain only letters, digits, spaces, or underscores")
	}
	return nil
}

func ValidatePreparationTime(value int32) error {
	if value <= 0 {
		return fmt.Errorf("preparation time must be positive")
	}
	return nil
}

func ValidateDifficultyLevel(value int32) error {
	if value <= 0 {
		return fmt.Errorf("difficulty level must be positive")
	}
	return nil
}

func ValidateCuisineType(value string) error {
	if !util.IsSupportedCuisineType(value) {
		return fmt.Errorf("cuisine type '%s' is not supported", value)
	}
	return nil
}

func ValidatePositiveInt32(value int32, fieldName string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be positive", fieldName)
	}
	return nil
}
