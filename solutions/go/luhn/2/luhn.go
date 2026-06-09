package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if id == "0" {
		return false
	}

	sum := 0
	for i := 0; i < len(id); i++ {
		if id[i] < '0' || id[i] > '9' {
			return false
		}

		distance := len(id) - 1 - i
		if distance%2 != 0 {
			b := 0
			a := int(id[i] - '0')
			if a*2 > 9 {
				b = a*2 - 9
			} else {
				b = a * 2
			}
			sum += b
		} else {
			b := int(id[i] - '0')
			sum += b
		}
	}

	return sum%10 == 0
}
