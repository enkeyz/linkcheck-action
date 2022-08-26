package link

import "testing"

func TestValidateURL(t *testing.T) {
	t.Run("returns no error, when the URL is valid", func(t *testing.T) {
		url := "https://github.com/practical-tutorials/project-based-learning"

		err := ValidateURL(url)
		assertNoError(t, err)
	})

	t.Run("returns error, when the URL is invalid", func(t *testing.T) {
		url := "https//github.com/practical-tutorials/project-based-learning"

		err := ValidateURL(url)
		assertError(t, err)
	})
}
