package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	if strings.Contains(isbn, "-") {
		isbn = strings.ReplaceAll(isbn, "-", "")
	}

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	lastDigit := isbn[len(isbn)-1:]

	for i := 0; i < 9; i++ {
		sum += int(isbn[i]-'0') * (10 - i)
	}

	if strings.ToUpper(lastDigit) == "X" {
		sum += 10
	} else {
		sum += int(isbn[9] - '0')
	}

	return sum%11 == 0
}
