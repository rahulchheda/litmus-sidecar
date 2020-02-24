package main

import (
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

func main() {
	handleSignal(signals.SetupSignalHandler())
}

func handleSignal(c <-chan struct{}) error {
	<-c
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("Time ", i)
		}
	}()
	fmt.Println("Received some signal")
	return nil
}
