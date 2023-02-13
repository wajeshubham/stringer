package stringer

import (
	"strconv"
)

func Reverse(str string) (result string) {
	for _, c := range str {
		result = string(c) + result
	}
	return result
}

func Inspect(str string, digits bool) (count int, kind string) {
	if !digits {
		return len(str), "char"
	}

	return inspectDigits(str), "digit"
}

func inspectDigits(str string) (count int) {
	for _, s := range str {
		_, err := strconv.Atoi(string(s))
		if err == nil {
			count++
		}
	}
	return count
}
