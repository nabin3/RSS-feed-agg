package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if len(authHeader) == 0 {
		return "", ErrNoAuthHeaderIncluded
	}

	splitedAuthHeader := strings.Split(authHeader, " ")
	if len(splitedAuthHeader) < 3 || splitedAuthHeader[1] != "ApiKey" {
		return "", errors.New("authorization header is malformed")
	}

	return splitedAuthHeader[2], nil
}
