package dtos

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}