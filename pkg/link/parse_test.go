package link

import (
	"testing"
)

func TestParseLink(t *testing.T) {
	t.Run("return the title and the url itself from the line", func(t *testing.T) {
		line := "- [Build an Interpreter](http://www.craftinginterpreters.com/) (Chapter 14 on is written in C)"

		want := &ParsedURL{
			Title: "Build an Interpreter",
			URL:   "http://www.craftinginterpreters.com/",
		}
		got, _ := ParseURL(line)

		if *got != *want {
			t.Errorf("got %#v, but want %#v", got, want)
		}
	})

	t.Run("return error if link and title couldn't be parsed", func(t *testing.T) {
		line := "- [Additional resources](#additional-resources)"

		_, err := ParseURL(line)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
