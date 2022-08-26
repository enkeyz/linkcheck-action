package link

import (
	"fmt"
	"regexp"
	"strings"
)

// https://regex101.com/r/r2cJgD/4
var markdownLinkRegex = regexp.MustCompile(`\[([^]]+)\]\((https?://[^\s/$.?#].[^\s]*)\)`)

type ParsedURL struct {
	Title, URL string
}

func ParseURL(line string) (*ParsedURL, error) {
	matches := markdownLinkRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return nil, fmt.Errorf("no url found in %q", line)
	}

	return &ParsedURL{
		Title: matches[0][1],
		URL:   strings.TrimSpace(matches[0][2]),
	}, nil
}
