package main

import (
	"fmt"

	"github.com/alrayyes/generictypealiasesdemo/oldpkg"
)

// This is possible since go 1.24.
type newUser[T any] = oldpkg.User[T]

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
