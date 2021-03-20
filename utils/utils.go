package utils

import (
	"regexp"
)

func IsHexColor(color string) bool {
	matched, err := regexp.MatchString(`^[0-9a-fA-F]{3,6}$`, color)
	if err != nil {
		return false
	}
	return matched
}
