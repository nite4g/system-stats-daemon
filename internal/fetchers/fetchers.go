package fetchers

import (
	"errors"
	"time"
)

type OsName string

var ErrorFetcher = errors.New("metric fetch error")

type MetricResult struct {
	Value     []string
	Error     error
	timestamp time.Time
}

type MetricCommand struct {
	name string
	os   OsName
	cmd  string
}

const (
	Macos OsName = "os"
	Linux OsName = "linux"
)
