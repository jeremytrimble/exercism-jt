// package twofer implements the ShareWith() function for sharing "one"'s with
// friends.
package twofer

import (
    "fmt"
)

// ShareWith returns the 
func ShareWith(s string) string {

    var other string

    if len(s) > 0 {
        other = s
    } else {
        other = "you"
    }

	return fmt.Sprintf("One for %s, one for me.", other)
}
