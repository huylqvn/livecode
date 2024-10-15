package livetest

import (
	"strings"
)

func revertString(str string) string {
	splitStr := strings.Split(str, " ")
	res := ""
	for i := 0; i < len(splitStr); i++ {
		res += revert(splitStr[i]) + " "
	}
	return res

}

func revert(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
