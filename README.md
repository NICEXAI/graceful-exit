# graceful-exit

Graceful exit by capturing program exit signals.Suitable for k8s pod logout、docker container stop、program exit and etc.

### Installation

Run the following command under your project:

> go get -u github.com/NICEXAI/graceful-exit

### Example

```go
package main

import (
	"fmt"
	gracefulExit "github.com/NICEXAI/graceful-exit"
)

func main() {
	// main process
	fmt.Println("Process Running...")

	graceful := gracefulExit.NewGracefulExit()

	graceful.RegistryHandle("exit", func() {
		// the specific logic of graceful exit
		fmt.Println("Process Stop...")
	})

	graceful.Capture()
}
```
