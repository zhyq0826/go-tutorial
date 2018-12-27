package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-cmd/cmd"
)

func main() {
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, "make", "-f", "/Users/zhaoyongqiang/workspace/devops/deploy/Makefile", "update_code_v12", "h=127.0.0.1")

	// Print STDOUT and STDERR lines streaming from Cmd
	go func() {
		for {
			select {
			case line := <-envCmd.Stdout:
				fmt.Println(line)
			case line := <-envCmd.Stderr:
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Cmd has finished but wait for goroutine to print all lines
	for len(envCmd.Stdout) > 0 || len(envCmd.Stderr) > 0 {
		time.Sleep(10 * time.Millisecond)
	}
}
