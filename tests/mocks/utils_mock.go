package mocks

import "errors"

var MockHashString = func(password string) (string, error) {
	if password == "error" {
		return "", errors.New("hashing failed")
	}
	return "hashedPassword", nil
}
