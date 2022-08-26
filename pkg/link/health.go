package link

import (
	"context"
	"fmt"
	"net/http"
)

func CheckHealth(ctx context.Context, url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `text/html`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36`)
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("checking link %s, got error %s", url, err)
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("checking link %s, got status code %d", url, res.StatusCode)
	}

	return nil
}
