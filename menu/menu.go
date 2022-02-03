package menu

import (
	"fmt"
	"time"
)

func Start() bool {

	var shutdown bool = false
	for !shutdown {
		var choice string

		fmt.Println("\n1. Dejarik\n2. Pazaak\n3. Sabacc\n4. quit" +
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
		case "4", "q", "Q", "quit", "Quit":
			fmt.Println("May the force be with you...")
			time.Sleep(1 * time.Second)
			fmt.Println("always")
			shutdown = true
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
	return shutdown
}
