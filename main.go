// Command generictypealiasesdemo builds the same generic user three ways —
// through an int32, a string and a local struct — and prints each one. The
// alias on line 10 is the thing being demonstrated; the rest is a way to watch
// it hold up against different type arguments.
package main

import (
	"fmt"

	"github.com/alrayyes/generictypealiasesdemo/user"
)

// This is possible since go 1.24.
type newUser[T any] = user.User[T]

func main() {
	intUser := newUser[int32]{
		ID:   1,
		Name: "Peter Integer",
	}
	stringUser := newUser[string]{
		ID:   "a",
		Name: "Peter String",
	}

	type customID struct {
		ID string
	}
	customUser := newUser[customID]{
		ID:   customID{ID: "a"},
		Name: "Peter Struct",
	}

	fmt.Println(intUser.AboutMe())
	fmt.Println(stringUser.AboutMe())
	fmt.Println(customUser.AboutMe())
}
