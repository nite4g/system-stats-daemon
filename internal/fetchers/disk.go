package fetchers

import (
	"os/exec"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func GetDiskSpace(os OsName) *MetricResult {

	res := &MetricResult{
		Value:     nil,
		Error:     nil,
		timestamp: time.Now(),
	}

	switch os {
	case Macos:
		// megabytes and mounted only
		out, err := exec.Command("/bin/df", "-lm").Output()
		if err != nil {
			log.Error().Err(err).Msg("os cmd error")
			res.Error = err
			return res
		}

		sOut := strings.Split(string(out), "\n")
		res.Value = sOut[1:]

	case Linux:
		// TODO: NOT IMPLEMENTED
		_, err := exec.Command("df", "-lm").Output()
		if err != nil {
			log.Error().Err(err).Msg("os cmd error")
			res.Error = err
			return res
		}

	default:
		log.Error().Err(ErrorFetcher).Str("os", string(os)).Msg("wrong Os")
		res.Error = ErrorFetcher
		return res
	}

	return res
}
