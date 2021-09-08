package main

import (
	"log"

	"github.com/n-arsenic/akunaki/migrate/cmd"
)

func main() {
	cmd := cmd.NewMigrateCmd()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
