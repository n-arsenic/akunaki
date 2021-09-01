package system

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Interrupt(stopper func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("server stopped")
		stopper()
		os.Exit(1)
	}()
}
