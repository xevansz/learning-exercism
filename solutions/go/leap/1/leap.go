// checks if it is a leap year
package leap

// checks if it is a leap year
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}

	return year%400 == 0
}
