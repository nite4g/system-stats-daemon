package fetchers

import (
	"os/exec"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// TODO: придумать как пробрасывать OsName более элегантно
func GetCPULA(os OsName) *MetricResult {

	res := &MetricResult{
		Value:     nil,
		Error:     nil,
		timestamp: time.Now(),
	}

	value, err := fetchCPULA(os)
	if err != nil {
		res.Error = err
		return res
	}

	// var metric float64

	switch os {
	case Macos:
		// value  "{ 2,02 2,11 1,99 }"
		s := strings.Fields(value)[1]
		s = strings.ReplaceAll(s, ",", ".")
		// metric, err = strconv.ParseFloat(s, 16)
		if err != nil {
			log.Error().Err(err).Msg("convertation from string error")
			res.Error = err
			return res
		}
		res.Value = []string{s}

	case Linux:
		// value "0.52 0.34 0.13 1/433 1769"
		s := strings.Fields(value)[0]
		// metric, err = strconv.ParseFloat(s, 16)
		if err != nil {
			log.Error().Err(err).Msg("convertation from string error")
			res.Error = err
			return res
		}
		res.Value = []string{s}

	default:
		log.Error().Err(ErrorFetcher).Str("os", string(os)).Msg("wrong Os")
		res.Error = ErrorFetcher
		return res
	}

	return res
}

func fetchCPULA(os OsName) (string, error) {

	// TODO: бесполезная функция, все равно есть ветвление в предыдущей
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
