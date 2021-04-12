package main

import (
	"context"
	"time"

	"github.com/nite4g/system-stats-daemon/internal/collector"
)

func main() {
	mc := collector.NewCollector(collector.Config{
		Name:     "cpu",
		Interval: 3 * time.Second,
	})
	ctx := context.Background()
	mc.Run(ctx)
}
