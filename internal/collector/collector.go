package collector

import (
	"sync"
	"time"
)

type Config struct {
	Name     string
	Interval time.Duration
}

type MetricCallback func() *MetricResult

type callback struct {
	name string
	cb   MetricCallback
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
	Run()
	AddCallBack(string, MetricCallback)
}

type collector struct {
	mu        sync.RWMutex
	opts      Config
	callbacks []callback
}

func (c *collector) AddCallBack(name string, cb MetricCallback) {
	c.mu.Lock()
	c.callbacks = append(c.callbacks, callback{
		name: name,
		cb:   cb,
	})
	c.mu.Unlock()
}

func (c *collector) Run() {
	var wg sync.WaitGroup
	result := MetricResponse
	for _, cb := range c.callbacks {
		wg.Add(1)
		go func(cb callback) {
			defer wg.Done()
			cb.cb()

		}(cb)
	}
	// да это не оптимально, но считаем что временем чтения метрики можно пренебречь
	wg.Wait()

	time.Sleep(c.opts.Interval * time.Second)
}

func NewCollector(opts Config) Collector {
	return &collector{
		opts: opts,
	}
}
