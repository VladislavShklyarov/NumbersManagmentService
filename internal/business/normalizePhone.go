package business

import (
	"errors"
	"regexp"
	"strings"
)

var nonDigits = regexp.MustCompile(`[^0-9+]`)

func NormalizePhone(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty phone")
	}

	clean := nonDigits.ReplaceAllString(input, "")

	if strings.HasPrefix(clean, "+") {
		digits := clean[1:]

		if len(digits) < 10 || len(digits) > 15 {
			return "", errors.New("invalid E.164 length")
		}

		return "+" + digits, nil
	}

	digits := strings.TrimPrefix(clean, "+")

	// RU cases
	if len(digits) == 11 && digits[0] == '8' {
		return "+7" + digits[1:], nil
	}

	if len(digits) == 10 {
		return "+7" + digits, nil
	}

	if len(digits) >= 11 && len(digits) <= 15 {
		return "+" + digits, nil
	}

	return "", errors.New("cannot normalize phone")
}
