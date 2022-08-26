package link

import (
	"context"
	"fmt"
	"net/http"
)

func CheckHealth(ctx context.Context, url string) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("checking link %s, got error %s", url, err)
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("checking link %s, got status code %d", url, res.StatusCode)
	}

	return nil
}
