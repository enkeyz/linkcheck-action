package config

import (
	"errors"
	"strconv"
	"time"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	ScannerTimeout     time.Duration
	ConcurrentRequests int
	FileName           string
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

	concRequestsStr := action.GetInput("concurrentRequests")
	if concRequestsStr == "" {
		return nil, errors.New("unable to get 'concurrentRequests' action input")
	}

	concRequests, err := strconv.Atoi(concRequestsStr)
	if err != nil {
		return nil, err
	}

	return &Config{
		ScannerTimeout:     timeout,
		ConcurrentRequests: concRequests,
		FileName:           "README.md",
	}, nil
}
