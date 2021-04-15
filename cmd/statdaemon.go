package main

import (
	"fmt"
	"time"

	"github.com/nite4g/system-stats-daemon/internal/collector"
	"github.com/nite4g/system-stats-daemon/internal/store"
)

func main() {
	storage := store.NewStorage()
	x, e := storage.Status()

	fmt.Printf("%v ### %v\n", x, e)
	mc := collector.NewCollector(collector.Config{
		Name:     "cpu",
		Interval: 3 * time.Second,
	})
	mc.AddCallBack("cpu_la", collector.MetricCallback(func() *collector.MetricResult {
		return collector.GetCPULA(collector.Macos)
	}))
	result := mc.Run()

	for name, r := range result {
		if r.Error != nil {
			fmt.Println(r.Error)
		}
		fmt.Printf("%v %v\n", name, r)
	}
}
