package user_test

import (
	"fmt"
	"github.com/alrayyes/generictypealiasesdemo/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_AboutMe(t *testing.T) {
	t.Run("test with an int ID", func(t *testing.T) {
		intUser := user.User[int]{
			ID:   0,
			Name: "Pete",
		}
		got := intUser.AboutMe()
		want := "Hi, my name is Pete and my ID is 0"

		assert.Equal(t, want, got)
	})

	t.Run("test with a string ID", func(t *testing.T) {
		stringUser := user.User[string]{
			ID:   "a",
			Name: "Pete",
		}
		got := stringUser.AboutMe()
		want := "Hi, my name is Pete and my ID is a"

		assert.Equal(t, want, got)
	})

	t.Run("test with a custom ID", func(t *testing.T) {
		type customID struct {
			ID string
		}

		customIDUser := user.User[customID]{
			ID:   customID{"a"},
			Name: "Pete",
		}
		got := customIDUser.AboutMe()
		want := "Hi, my name is Pete and my ID is {a}"

		assert.Equal(t, want, got)
	})
}

func ExampleUser_AboutMe() {
	intUser := user.User[int]{
		ID:   1,
		Name: "James",
	}
	fmt.Println(intUser.AboutMe())
	// Output: Hi, my name is James and my ID is 1
}
