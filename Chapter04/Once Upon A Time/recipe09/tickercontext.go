package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	var wg sync.WaitGroup
	ctx, cancelCtx := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-ctx.Done():
				fmt.Println("Goroutine closing")
				return
			}
		}
	}()

	// Block until the signal is received
	<-c

	// Stop the goroutine by calling the returned cancel function for the context.
	// Multiple GO-routines can use the same ctx.Done channel to trigger a graceful shutdown.
	cancelCtx()

	// Wait until the GO-routine have stopped
	wg.Wait()
	fmt.Println("Application stopped")
}
