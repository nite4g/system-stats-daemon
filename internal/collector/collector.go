package collector

import (
	"context"
	"fmt"
	"time"
)

type Config struct {
	Name     string
	Interval time.Duration
}

type MetricValue struct {
	// metric name i.e. "cpu_la", "mem_free", "disk_space"
	name string
	// current value
	value float32
	// unit i.e. "MB", "inodes", "secs"
	unit string
}

type Collector interface {
	Run(context.Context)
}

type collector struct {
	opts Config
}

func (c *collector) Run(ctx context.Context) {
	fmt.Printf("Name: %v, Duration: %v", c.opts.name, c.opts.interval)
}

func NewCollector(opts Config) Collector {
	return &collector{
		opts,
	}
}
