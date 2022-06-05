package main

import (
	"context"
	// standard
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// internal
	"MicroservicesWithGo/handlers"
)

const PORT = ":2022"

func main() {
	var logs = log.New(os.Stdout, "MicroserviceWithGo", log.LstdFlags)
	var helloHandler = handlers.NewHello(logs)
	var productsHandler = handlers.NewProduct(logs)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/products", productsHandler)

	server := &http.Server{
		Addr:         PORT,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logs.Fatal(err)
			return
		}
	}()

	signalChain := make(chan os.Signal)
	signal.Notify(signalChain, os.Interrupt)
	signal.Notify(signalChain, os.Kill)
	sg := <-signalChain
	logs.Println("Terminated Successfully", sg)
	timeContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := server.Shutdown(timeContext)
	if err != nil {
		logs.Println(err)
		return
	}
}
