package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Koratsama/StarWarsClassics/menu"
)

//main method of star wars classics
func main() {

	absPath, _ := filepath.Abs("../StarWarsClassics/ascii/StarWarsClassic.txt")
	content, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println("\nWelcome to Star Wars Classics!")

	var shutdown bool = false

	for !shutdown {
		shutdown = menu.Start()
	}
	//shutdown = menu.Start()

	if shutdown {
		go gracefulShutdown()
		forever := make(chan int)
		<-forever
	}
}

func gracefulShutdown() {
	//s := make(chan os.Signal, 1)
	//signal.Notify(s, os.Interrupt)
	//signal.Notify(s, syscall.SIGTERM)
	os.Interrupt.Signal()
	go func() {
		//<-s
		fmt.Println("Sutting down gracefully...")
		// clean up here
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}()
}
