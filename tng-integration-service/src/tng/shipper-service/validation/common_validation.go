package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
	"strings"
	"time"
	"tng/common/models"
)

const (
	userNamePattern    = "^[A-Za-z0-9]*$"
	upperString        = "^(.*?[A-Z])"
	lowerString        = "^(.*?[a-z])"
	specialCharacter   = `!@#$%^&*()-_=+\|[]{};:/?.><`
	validLogoImage     = `([a-zA-Z0-9\s_\\.\-\(\):])+(.png|.jpg|.jpeg)$`
	phoneNumberPattern = `((^84)+([0-9]{9,11}))|((^0)+([0-9]{9}))$`
)

var (
	userNameRegex       = regexp.MustCompile(userNamePattern)
	upperStringRegex    = regexp.MustCompile(upperString)
	lowerStringRegex    = regexp.MustCompile(lowerString)
	validLogoImageRegex = regexp.MustCompile(validLogoImage)
	phoneNumberRegex    = regexp.MustCompile(phoneNumberPattern)
)

// Match will check the value of fl matching with regexp2 or not.
func Match(reg *regexp.Regexp, fl validator.FieldLevel) bool {
	return reg.MatchString(fl.Field().String())
}

// Iso8601DateTime Validate if string is match with date time format ISO 8601
func Iso8601DateTime(fl validator.FieldLevel) bool {
	_, err := time.Parse(models.FormatDateTime, fl.Field().String())
	return err == nil
}

// UserNamePattern will check a string not containing special characters and whitespace and Vietnamese characters.
func UserNamePattern(fl validator.FieldLevel) bool {
	return Match(userNameRegex, fl)
}

// ValidatePassword will only accept a string containing numbers 8 characters, UpCharacter, LowerCharacter, Special Character
func ValidatePassword(fl validator.FieldLevel) bool {
	return Match(upperStringRegex, fl) && Match(lowerStringRegex, fl) && strings.ContainsAny(fl.Field().String(), specialCharacter)
}

// ValidLogoImage will only accept a string that end with .png, .jpeg, .jpg
func ValidLogoImage(fl validator.FieldLevel) bool {
	return Match(validLogoImageRegex, fl)
}

// ValidPhoneNumber will only accept a string is phone number format
func ValidPhoneNumber(fl validator.FieldLevel) bool {
	return Match(phoneNumberRegex, fl)
}
