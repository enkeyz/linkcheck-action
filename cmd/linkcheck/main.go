package main

import (
	"log"

	"github.com/enkeyz/go-linkcheck/internal/linkscanner"
	"github.com/enkeyz/go-linkcheck/pkg/config"
	"github.com/sethvargo/go-githubactions"
)

func main() {
	action := githubactions.New()
	cfg, err := config.NewFromInputs(action)
	if err != nil {
		log.Fatal(err)
	}

	linkScanner := linkscanner.New(cfg)
	linkScanner.Scan()
}
