package main

import (
	"log"

	"github.com/n-arsenic/akunaki/server/cmd"
)

func main() {
	cmd := cmd.NewRootCmd()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
