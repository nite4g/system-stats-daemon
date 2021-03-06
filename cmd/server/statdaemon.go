package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/nite4g/system-stats-daemon/internal/statistic"

	"github.com/nite4g/system-stats-daemon/internal/collector"
	"github.com/nite4g/system-stats-daemon/internal/fetchers"
	"github.com/nite4g/system-stats-daemon/internal/store"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const fetchInterval = 3

func main() {
	storage := store.NewStorage()
	x, e := storage.Status()

	fmt.Printf("%v ### %v\n", x, e)
	mc := collector.NewCollector(collector.Config{
		Name: "Main",
	})

	mc.AddCallBack("cpu_la", collector.MetricCallback(func() *fetchers.MetricResult {
		return fetchers.GetCPULA(fetchers.Macos)
	}))

	mc.AddCallBack("disks_space", collector.MetricCallback(func() *fetchers.MetricResult {
		return fetchers.GetDiskSpace(fetchers.Macos)
	}))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			result := mc.Process()
			for name, r := range result {
				if r.Error != nil {
					fmt.Println(r.Error)
				}
				fmt.Printf("%#v %#v\n", name, r)
			}
			time.Sleep(fetchInterval * time.Second)
		}
	}()

	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
	}
	s := statistic.Server{}
	grpcServer := grpc.NewServer()
	statistic.RegisterStatisticServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Error().Err(err).Msg("Error running grpc server")
	}

	wg.Wait()
}
