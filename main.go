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
	fmt.Println("\nWelcome to Star Wars Classics!")
	for !shutdown {
		var choice string

		fmt.Println("\n1. Dejarik\n2. Pazaak\n3. Sabacc" +
			"\nPlease select a game to launch:")

		fmt.Scanf("%s\n", &choice)

		switch choice {
		case "1", "Dejarik", "dejarik":
			fmt.Println("\nThank you for choosing Dejarik!" +
				"\n... unfortunately this game is not available at this time." +
				"\nPlease choose another.")
		case "2", "Pazaak", "pazaak":
			fmt.Println("\nThank you for choosing Pazaak!" +
				"\n... unfortunately this game is not available at this time." +
				"\nPlease choose another.")
		case "3", "Sabacc", "sabacc":
			fmt.Println("\nThank you for choosing Sabacc!" +
				"\n... unfortunately this game is not available at this time." +
				"\nPlease choose another.")
		case "q", "quit":
			fmt.Println("May the force be with you...")
			time.Sleep(1 * time.Second)
			fmt.Println("always")
			shutdown = true
		default:
			fmt.Println("Invalid option. Please choose again.")
		}

	}

	os.Exit(0)

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
