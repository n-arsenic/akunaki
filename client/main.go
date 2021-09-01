package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/n-arsenic/grpc-service/client/generated/proto/echo"

	"google.golang.org/grpc"
)

type Config struct {
	UserName string
	Addr     string
}

func main() {
	var config Config
	flag.StringVar(&config.UserName, "client", "", "client name")
	flag.StringVar(&config.Addr, "address", ":8080", "server address")
	flag.Parse()
	if config.UserName == "" || config.Addr == "" {
		log.Print("bad parameters")
		os.Exit(2)
	}

	conn, err := grpc.Dial(config.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Wellcome(ctx, &pb.Request{User: config.UserName})
	if err != nil {
		log.Fatalf("could not process request: %v", err)
	}
	log.Printf(resp.GetMessage())
}
