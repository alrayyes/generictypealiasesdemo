package oldpkg_test

import (
	"github.com/alrayyes/generictypealiasesdemo/oldpkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_AboutMe(t *testing.T) {
	t.Run("test with an int ID", func(t *testing.T) {
		user := oldpkg.User[int]{
			ID:   0,
			Name: "Pete",
		}
		got := user.AboutMe()
		want := "Hi, my name is Pete and my ID is 0"

		assert.Equal(t, want, got)
	})

	t.Run("test with a string ID", func(t *testing.T) {
		user := oldpkg.User[string]{
			ID:   "a",
			Name: "Pete",
		}
		got := user.AboutMe()
		want := "Hi, my name is Pete and my ID is a"

		assert.Equal(t, want, got)
	})

	t.Run("test with a custom ID", func(t *testing.T) {
		type customID struct {
			ID string
		}

		user := oldpkg.User[customID]{
			ID:   customID{"a"},
			Name: "Pete",
		}
		got := user.AboutMe()
		want := "Hi, my name is Pete and my ID is {a}"

		assert.Equal(t, want, got)
	})
}
