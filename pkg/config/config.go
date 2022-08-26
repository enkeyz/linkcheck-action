package config

import "time"

type Config struct {
	ScannerTimeout time.Duration
	FileName       string
}
