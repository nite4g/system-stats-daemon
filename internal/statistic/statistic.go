package statistic

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) Send(ctx context.Context, in *StatisticMessage) (*StatisticMessage, error) {
	log.Printf("Receive message body from client: %s", in.Metric)
	return &StatisticMessage{Value: "Hello From the Server!"}, nil
}
