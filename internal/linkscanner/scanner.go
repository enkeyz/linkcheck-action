package linkscanner

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/enkeyz/go-linkcheck/pkg/config"
	"github.com/enkeyz/go-linkcheck/pkg/link"
)

// check if README.md exists
// read every line into a string slice
// parse each line for title and link
// TODO check if they valid urls
// run health check on each url and log the result

type LinkScanner struct {
	timeout  time.Duration
	fileName string
}

func New(cfg *config.Config) *LinkScanner {
	return &LinkScanner{
		timeout:  cfg.ScannerTimeout,
		fileName: cfg.FileName,
	}
}

func (l *LinkScanner) Scan() {

	readme, err := l.openReadme()
	if err != nil {
		log.Fatal(err)
	}
	defer readme.Close()

	lines := l.scanFile(readme)
	if len(lines) == 0 {
		log.Fatalf("No text found in %s", l.fileName)
	}

	parsedURLs := l.parseURLs(lines)
	if len(parsedURLs) == 0 {
		log.Fatalf("No links found in %s", l.fileName)
	}
	log.Printf("Found %d links", len(parsedURLs))

	l.doHealthCheck(parsedURLs)
}

func (l *LinkScanner) openReadme() (*os.File, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	readme, err := os.Open(filepath.Join(cwd, l.fileName))
	if err != nil {
		return nil, err
	}

	return readme, nil
}

func (l *LinkScanner) scanFile(file *os.File) []string {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lines := make([]string, 0)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			continue
		}

		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func (l *LinkScanner) parseURLs(lines []string) []*link.ParsedURL {
	parsedURLs := make([]*link.ParsedURL, 0)
	for _, line := range lines {
		parsedURL, err := link.ParseURL(line)
		if err != nil {
			continue
		}

		if err = link.ValidateURL(parsedURL.URL); err != nil {
			continue
		}

		parsedURLs = append(parsedURLs, parsedURL)
	}

	return parsedURLs
}

func (l *LinkScanner) doHealthCheck(parsedURLs []*link.ParsedURL) {
	errorChan := make(chan error)
	for _, parsedURL := range parsedURLs {
		go func(url string) {
			cancellingCtx, cancel := context.WithCancel(context.Background())
			time.AfterFunc(l.timeout, cancel)

			errorChan <- link.CheckHealth(cancellingCtx, url)
		}(parsedURL.URL)
	}

	for i := 0; i < len(parsedURLs); i++ {
		err := <-errorChan
		if err != nil {
			log.Printf("Error: %s", err)
		}
	}
}
