package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func readYesNo() bool {
	b := make([]byte, 1)

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	// Wait reply
	for {
		os.Stdin.Read(b)
		character := string(b)
		if character == "y" || character == "Y" || character == "o" || character == "O" {
			return true
		}
		if character == "n" || character == "N" {
			return false
		}
		fmt.Print("?")
	}
}
