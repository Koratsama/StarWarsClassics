package menu_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Koratsama/StarWarsClassics/menu"
)

func TestMenuSabacc(t *testing.T) {

	content := []byte("Sabacc\r\n")
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

	if shutdown == true {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMenuCorellianSpike(t *testing.T) {
	//skip this test for now since we are starting to implement round logic.
	t.Skip()
	content := []byte("Corellian Spike\r\n")
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

	if shutdown == true {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMenuCoruscantShift(t *testing.T) {

	content := []byte("Coruscant Shift\r\n")
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

	if shutdown == true {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMenuQuitPath(t *testing.T) {
	content := []byte("Quit\r\n")
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

func TestMenuInvalidInput(t *testing.T) {
	content := []byte("Dejarik\r\n")
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

	if shutdown == true {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMenuError(t *testing.T) {
	content := []byte("Dejarik\r") // no '\n' character should throw error.
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
		fmt.Println("Fail! Error expected")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
