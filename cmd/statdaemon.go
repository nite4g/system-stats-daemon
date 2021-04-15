package main

import (
	"fmt"
	"time"

	"github.com/nite4g/system-stats-daemon/internal/collector"
)

func main() {
	mc := collector.NewCollector(collector.Config{
		Name:     "cpu",
		Interval: 3 * time.Second,
	})
	// ctx := context.Background()
	mc.AddCallBack("cpu_la", collector.MetricCallback(func() *collector.MetricResult {
		return collector.GetCpuLA(collector.Macos)
	}))
	result := mc.Run()

	for _, r := range result {
		if r.Error != nil {
			fmt.Println(r.Error)
		}
		fmt.Printf("%v\n", r)
	}

}
