package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi"
	"github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/server"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Starting bukakoneksi...")
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/bukalapak/hackathon_may2018_bukakoneksi/bukakoneksi/.env")

	dsn := os.Getenv("DEVELOPMENT_DATABASE_URL")
	bukakoneksi.InitDB(dsn)

	router := server.Router()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Bukakoneksi started @:8062")
		http.ListenAndServe(":8062", router)
	}()

	<-sigChan
	fmt.Println()
	log.Println("Shutting down bukakoneksi...")
	log.Println("Bukakoneksi stopped")
}
