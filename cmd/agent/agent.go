package main

import (
	"log"

	"github.com/nite4g/system-stats-daemon/internal/statistic"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := statistic.NewStatisticServiceClient(conn)

	response, err := c.Send(context.Background(), &statistic.StatisticMessage{Value: "Hello From the Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Value)

}
