package twofer

import "fmt"

// ShareWith func
// Given a name, return a string with the message:
//   One for X, one for me.
// where x is given name
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
