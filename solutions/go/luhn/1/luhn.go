package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if id == "0" {
		return false
	}

	sum := 0
	bs := []byte(id)

	for i := len(id) - 1 - 1; i >= 0; i -= 2 {

		b := 0
		a := int(id[i] - '0')
		if a*2 > 9 {
			b = a*2 - 9
		} else {
			b = a * 2
		}
		bs[i] = byte(b) + '0'
	}

	for i := len(bs) - 1; i >= 0; i-- {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
		sum += int(bs[i] - '0')
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}
