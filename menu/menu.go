package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Koratsama/StarWarsClassics/sabacc"
)

/*
Name: Start
Purpose: Prompts the user a menu of games to choose from.
Parameters: None
*/
func Start() bool {

	var shutdown bool = false
	var choice string

	fmt.Println("\n1. Sabacc\n2. Corellian Spike\n3. Coruscant Shift\n4. Quit" +
		"\nPlease select a game to launch:")

	in := bufio.NewReader(os.Stdin)
	choice, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return true
	}
	//TODO: see if there is a better solution for reading input for mac&windows.
	choice = strings.Trim(choice, "\n") //for reading in mac input
	choice = strings.Trim(choice, "\r") //for reading in windows input.

	switch choice {
	case "1", "Sabacc", "sabacc":
		fmt.Println("\nThank you for choosing Sabacc!" +
			"\n... unfortunately this game is not available at this time." +
			"\nPlease choose another.")
	case "2", "Corellian Spike", "corellian spike":
		fmt.Println("\nThank you for choosing Corellian Spike Sabacc!")
		sabacc.Start()
		fmt.Println("\nThank you for playing!...")
	case "3", "Coruscant Shift", "coruscant shift":
		fmt.Println("\nThank you for choosing Coruscant Shift Sabacc!" +
			"\n... unfortunately this game is not available at this time." +
			"\nPlease choose another.")
	case "4", "q", "Q", "Quit", "quit":
		fmt.Println("May the force be with you...")
		time.Sleep(1 * time.Second)
		fmt.Println("always")
		shutdown = true
	default:
		fmt.Printf("\nInvalid option: %s\nPlease choose again.\n", choice)
	}
	return shutdown
}
