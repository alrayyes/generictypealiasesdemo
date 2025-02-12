package oldpkg

import "fmt"

type User[T any] struct {
	ID   T
	Name string
}

func (u User[T]) AboutMe() string {
	return fmt.Sprintf("Hi, my name is %s and my ID is %v", u.Name, u.ID)
}
