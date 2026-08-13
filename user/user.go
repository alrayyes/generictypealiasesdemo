// Package user holds the generic type the demo aliases. It is deliberately
// small: the interesting part is that User's type parameter survives being
// renamed through an alias in another package, not anything User itself does.
package user

import "fmt"

// User contains everything we could possibly want to know about a human being.
type User[T any] struct {
	ID   T
	Name string
}

// AboutMe tells the world what the users name & id is.
func (u User[T]) AboutMe() string {
	return fmt.Sprintf("Hi, my name is %s and my ID is %v", u.Name, u.ID)
}
