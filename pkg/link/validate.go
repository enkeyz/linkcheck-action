package link

import (
	"fmt"
	"net/url"
)

func ValidateURL(rawUrl string) error {
	if _, err := url.ParseRequestURI(rawUrl); err != nil {
		return fmt.Errorf("error validating url %s: %s", rawUrl, err)
	}

	return nil
}
