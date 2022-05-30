package menu_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Koratsama/StarWarsClassics/menu"
)

func TestMenuStart(t *testing.T) {

	//TODO: actually test user input for the menu
	fmt.Printf("test Okay!")
}

func TestMenuQuitPath(t *testing.T) {
	content := []byte("Quit\n")
	tmpfile, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	//run menu
	var shutdown bool = menu.Start()

	if shutdown != true {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
