package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var sigs = make(chan os.Signal)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		//sig := <-sigs
		fmt.Println("shutdown ", <-sigs)
		close(done)
	}()
	<-done
	fmt.Println("Graceful")
}
