package strutil

import (
	"encoding/base64"
	"math/rand"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type set struct {
}

// Contains implementation of runes.Set.
func (s *set) Contains(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

var defaultSet = &set{}

// NormalizeUnicode normalizes a unicode string.
func NormalizeUnicode(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(defaultSet), norm.NFC)
	s, _, _ = transform.String(t, s)
	s = strings.Replace(s, "Đ", "D", -1)
	s = strings.Replace(s, "đ", "d", -1)
	return s
}

// DecodeBase64 will decode msg to string
func DecodeBase64(msg string) (string, error) {
	if msg == "" {
		return "", nil
	}

	decode, err := base64.StdEncoding.DecodeString(msg)

	if err != nil {
		return "", err
	}

	return string(decode), nil
}

// UniqueInt64 remove duplicate item
func UniqueInt64(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// define chars random
const charsPattern = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+"

// RandPassword random password
func RandPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charsPattern[rand.Intn(len(charsPattern))]
	}
	return string(b)
}

// Check if string is empty
func IsEmpty(str string) bool {
	return len(str) == 0
}

// Check if string is not empty
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}