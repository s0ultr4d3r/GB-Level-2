package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var workers = make(chan struct{}, 1)

func main() {

	// пул с буфером в 1
	for i := 1; i <= 1001; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers
			}()
			fmt.Println(job)
		}(i)

	}

	go sysSigns()
}

func sysSigns() {
	log.Print("start")
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	sig := <-signalChan

	for {

		timeout := (time.After(1 * time.Second))
		select {
		default:
			log.Printf("got %s signal", sig.String())
			log.Print("end")

		case <-timeout:
			return
		}
	}
}
