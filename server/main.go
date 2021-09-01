package main

import (
	"log"

	"github.com/n-arsenic/grpc-service/server/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
