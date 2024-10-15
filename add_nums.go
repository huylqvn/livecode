package livetest

import "strings"

func addStrings(num1 string, num2 string) string {
	var res strings.Builder

	l1, l2 := len(num1)-1, len(num2)-1
	carry := 0
	for l1 >= 0 || l2 >= 0 {
		sum := carry
		if l1 >= 0 {
			sum += int(num1[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			sum += int(num2[l2] - '0')
			l2--
		}
		res.WriteByte(byte(sum%10) + '0')
		carry = sum / 10
	}

	if carry > 0 {
		res.WriteByte(byte(carry) + '0')
	}

	runes := []rune(res.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str := string(runes)
	return str
}
