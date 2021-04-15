package collector

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type OsName string

var FetchError = errors.New("Metric fetch error.")

type MetricResult struct {
	Value     float64
	Error     error
	timestamp time.Time
}

const (
	Macos OsName = "os"
	Linux OsName = "linux"
)

// TODO: придумать как пробрасывать OsName более элегантно

func GetCpuLA(os OsName) *MetricResult {
	value, err := fetchCpuLA(os)

	if err != nil {
		return &MetricResult{Value: 0, Error: err}
	}

	var metric float64

	switch os {
	case Macos:
		// value  "{ 2,02 2,11 1,99 }"
		s := strings.Fields(value)[1]
		metric, err = strconv.ParseFloat(s, 32)
		if err != nil {

			return &MetricResult{Value: 0, Error: FetchError}
		}
	case Linux:
		// value "0.52 0.34 0.13 1/433 1769"
		s := strings.Fields(value)[0]
		metric, err = strconv.ParseFloat(s, 32)
		if err != nil {
			return &MetricResult{Value: 0, Error: FetchError}
		}

	default:
		return &MetricResult{Value: 0, Error: FetchError}
	}

	return &MetricResult{Value: metric, Error: nil}
}

func fetchCpuLA(os OsName) (string, error) {
	// sysctl -n vm.loadavg  for Mac OS
	var cmd string

	switch os {
	case Macos:
		cmd = "/usr/sbin/sysctl -n vm.loadavg"
	case Linux:
		cmd = "/usr/bin/cat /proc/loadavg"
	default:
		return "", FetchError
	}

	app := exec.Command(cmd)
	out, err := app.Output()
	if err != nil {
		return "", FetchError
	}

	return string(out), nil
}
