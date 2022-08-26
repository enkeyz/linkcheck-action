package config

import (
	"errors"
	"time"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	ScannerTimeout time.Duration
	FileName       string
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {
	timeoutStr := action.GetInput("timeout")
	if timeoutStr == "" {
		return nil, errors.New("unable to get 'timeout' action input")
	}

	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		ScannerTimeout: timeout,
		FileName:       "README.md",
	}, nil
}
