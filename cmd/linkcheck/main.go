package main

import (
	"time"

	"github.com/enkeyz/go-linkcheck/internal/linkscanner"
	"github.com/enkeyz/go-linkcheck/pkg/config"
)

func main() {
	linkScanner := linkscanner.New(&config.Config{ScannerTimeout: 10 * time.Second, FileName: "README.md"})
	linkScanner.Scan()
}
