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

type Collector interface {
	Run() map[string]MetricResult
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

func (c *collector) Run() map[string]MetricResult {
	var wg sync.WaitGroup

	result := map[string]MetricResult{}

	for _, cb := range c.callbacks {
		wg.Add(1)
		go func(cb callback) {
			defer wg.Done()
			c.mu.Lock()
			result[cb.name] = *cb.cb()
			c.mu.Unlock()
		}(cb)
	}
	// это не оптимально, но считаем что временем чтения метрики можно пренебречь
	wg.Wait()

	return result
}

func NewCollector(opts Config) Collector {
	return &collector{
		opts: opts,
	}
}
