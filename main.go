package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//main method of star wars classics
func main() {

	var shutdown bool = false

	for !shutdown {

		time.Sleep(10 * time.Second)
		shutdown = true
	}

	go gracefulShutdown()
	forever := make(chan int)
	<-forever
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Sutting down gracefully...")
		// clean up here
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()
}
