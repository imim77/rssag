package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(value, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of  authentication header")
	}
	return vals[1], nil
}
