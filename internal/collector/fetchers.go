package collector

import (
	"os/exec"
	"strconv"
	"strings"
)

type OsName string

var FetchError = error.New("Metric fetch error")

const (
	macos OsName = "os"
	linux OsName = "linux"
)

// TODO: придумать как пробрасывать OsName более элегантное решение

func GetCpuLA(os OsName) (float64, error) {
	value, err := fetchCpuLA(os)
	if err != nil {
		return 0, FetchError
	}

	var metric float64

	switch os {
	case macos:
		// value  "{ 2,02 2,11 1,99 }"
		s := strings.Fields(value)[1]
		metric, err = strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, FetchError
		}
	case linux:
		// value "0.52 0.34 0.13 1/433 1769"
		s := strings.Fields(value)[0]
		metric, err = strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, FetchError
		}

	default:
		return 0, FetchError
	}

	return metric, nil
}

func fetchCpuLA(os OsName) (string, error) {
	// sysctl -n vm.loadavg  for Mac OS
	var cmd string

	switch os {
	case macos:
		cmd = "/usr/sbin/sysctl -n vm.loadavg"
	case linux:
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
