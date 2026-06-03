// It's package that determines wat you will say as you give away the extra cookie
package twofer

import "fmt"

// func that implements it
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}

	return "One for you, one for me."
}
