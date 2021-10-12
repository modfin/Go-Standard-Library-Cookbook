package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	prc := exec.Command("./inputrequired.sh")
	prc.Stdout = os.Stdout
	prc.Stdin = os.Stdin
	// note @jonas: at least on macOS "read" prompt (-p) seems to output to stdErr
	// so we need to pipe that too
	prc.Stderr = os.Stderr

	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process ran successfully")
	}
}
