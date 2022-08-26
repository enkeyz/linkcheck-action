package link

import (
	"context"
	"fmt"
	"net/http"
)

func CheckHealth(ctx context.Context, url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("checking link %s, got error %s", url, err)
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("checking link %s, got status code %d", url, res.StatusCode)
	}

	return nil
}
