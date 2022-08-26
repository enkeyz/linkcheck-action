package config

import (
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
	timeoutStr := githubactions.GetInput("timeout")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, err
	}

	concRequestsStr := githubactions.GetInput("concurrentRequests")
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
