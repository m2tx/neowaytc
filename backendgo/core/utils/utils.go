package utils

import "strconv"

func AllDigitsEquals(number string) bool {
	n := number[0]
	for i := 1; i < len(number); i++ {
		if n != number[i] {
			return false
		}
	}
	return true
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ToInt(n byte) int {
	return int(n - '0')
}

func StringToInt(n string) int {
	nb, err := strconv.Atoi(n)
	if err != nil {
		return 0
	}
	return nb
}
