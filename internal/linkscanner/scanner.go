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
	"golang.org/x/sync/semaphore"
)

type LinkScanner struct {
	timeout            time.Duration
	concurrentRequests int
	fileName           string
}

func New(cfg *config.Config) *LinkScanner {
	return &LinkScanner{
		timeout:            cfg.ScannerTimeout,
		concurrentRequests: cfg.ConcurrentRequests,
		fileName:           cfg.FileName,
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
	sem := semaphore.NewWeighted(int64(l.concurrentRequests))

	for _, parsedURL := range parsedURLs {
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatal("Unable to acquire semaphore")
		}

		go func(url string) {
			cancellingCtx, cancel := context.WithCancel(context.Background())
			time.AfterFunc(l.timeout, cancel)

			errorChan <- link.CheckHealth(cancellingCtx, url)
			sem.Release(1)
		}(parsedURL.URL)
	}

	for i := 0; i < len(parsedURLs); i++ {
		err := <-errorChan
		if err != nil {
			log.Printf("Error: %s", err)
		}
	}
}
