package api

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosa12978/Rankee/internal/api/routers"
)

func Run() {
	// ? Maybe move to another function
	rand.Seed(time.Now().Unix())

	router := routers.NewMainRouter()

	server := http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8987",
	}
	go server.ListenAndServe()

	log.Printf("Server started")
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("%s", (<-sign).String())
}
