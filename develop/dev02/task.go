package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func UnpackString(s string) (string, error) {
	var result strings.Builder
	var repeatCount int
	var escaped bool

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) && !escaped {
			count, err := strconv.Atoi(string(s[i]))
			if err != nil {
				return "", err
			}
			repeatCount = repeatCount*10 + count
		} else if s[i] == '\\' && !escaped {
			escaped = true
		} else {
			if repeatCount == 0 {
				repeatCount = 1
			}
			result.WriteString(strings.Repeat(string(s[i]), repeatCount))
			repeatCount = 0
			escaped = false
		}
	}

	if escaped {
		return "", errors.New("некорректная строка")
	}

	return result.String(), nil
}
