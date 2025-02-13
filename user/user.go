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
