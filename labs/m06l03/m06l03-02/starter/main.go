package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	fmt.Println("ready")
	<-stop
	fmt.Println("stopping")
}
