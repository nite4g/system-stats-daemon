package collector

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type OsName string

var ErrorFetcher = errors.New("Metric fetch error.")

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
	value, err := fetchCPULA(os)

	res := &MetricResult{
		Value:     0,
		Error:     nil,
		timestamp: time.Now(),
	}

	if err != nil {
		res.Error = err
		return res
	}

	var metric float64

	switch os {
	case Macos:
		// value  "{ 2,02 2,11 1,99 }"
		s := strings.Fields(value)[1]
		s = strings.ReplaceAll(s, ",", ".")
		metric, err = strconv.ParseFloat(s, 16)
		if err != nil {
			log.Error().Err(err).Msg("convertation from string error")
			res.Error = err
			return res
		}
		res.Value = metric

	case Linux:
		// value "0.52 0.34 0.13 1/433 1769"
		s := strings.Fields(value)[0]
		metric, err = strconv.ParseFloat(s, 16)
		if err != nil {
			log.Error().Err(err).Msg("convertation from string error")
			res.Error = err
			return res
		}
		res.Value = metric

	default:
		log.Error().Err(ErrorFetcher).Str("os", string(os)).Msg("wrong Os")
		res.Error = ErrorFetcher
		return res
	}

	return res
}

func fetchCPULA(os OsName) (string, error) {
	// sysctl -n vm.loadavg  for Mac OS
	var cmd *exec.Cmd

	switch os {
	case Macos:
		cmd = exec.Command("/usr/sbin/sysctl", "-n", "vm.loadavg")
	case Linux:
		cmd = exec.Command("/usr/bin/cat", "/proc/loadavg")
	default:
		return "", ErrorFetcher
	}

	out, err := cmd.Output()
	if err != nil {
		log.Error().Err(err).Msg("os cmd error")
		return "", err
	}

	return string(out), nil
}
