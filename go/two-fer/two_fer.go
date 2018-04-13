package twofer

import "fmt"

// ShareWith ...
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %v, one for me.", name)
	}
	return "One for you, one for me."
}
