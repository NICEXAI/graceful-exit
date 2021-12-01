package main

import (
	"fmt"
	gracefulexit "github.com/NICEXAI/graceful-exit"
)

func main() {
	// main process
	fmt.Println("Process Running...")

	graceful := gracefulexit.NewGracefulExit()

	graceful.RegistryHandle("exit", func() {
		// the specific logic of graceful exit
		fmt.Println("Process Stop...")
	})

	graceful.Capture()
}
