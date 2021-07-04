package utils

import (
	"errors"
	"os"
	"strconv"
)

var ErrEmptyEnvVar = errors.New("empty environment variable")

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEmptyEnvVar
	}
	return v, nil
}

func GetEnvBool(key string) (bool, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func GetEnvInt(key string) (int, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
