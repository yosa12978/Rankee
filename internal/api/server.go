package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log.Printf("Server started")
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("%s", (<-sign).String())
}
