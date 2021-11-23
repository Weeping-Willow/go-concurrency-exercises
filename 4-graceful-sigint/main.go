//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Create a process
	proc := MockProcess{}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		killCommandCount := 0
		for sig := range c {
			fmt.Println(killCommandCount)
			if sig != os.Interrupt {
				continue
			}

			if killCommandCount == 0 {
				killCommandCount++
				go proc.Stop()
			} else if killCommandCount == 1 {
				fmt.Println(1)
				os.Exit(1)
			}
		}
	}()

	// Run the process (blocking)
	proc.Run()
}
