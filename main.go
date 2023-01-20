package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/light-state-machine/runtime"
)

func main() {
	fmt.Println("hello , I`m coming. Demo for K8s Operator")
	rt := runtime.NewSMRuntime()
	rt.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop
	rt.Shutdown()
}
