package utils

import (
	"math/rand"
	"net/mail"
	"regexp"
)

func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsCellphone(cellphone string) bool {
	reg := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return reg.MatchString(cellphone)
}
