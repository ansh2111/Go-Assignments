package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"context"
	"log"
	"net/http"

	"empapp"
)


func main() {
	ctx:= context.Background()
	errChan := make(chan error)

    	go func() {
       		c := make(chan os.Signal, 1)
        	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
        	errChan <- fmt.Errorf("%s", <-c)
    	}()

	go func() {
        	log.Println("empapp is listening on port:", ":8080")
        	handler := empapp.NewHTTPServer(ctx)
        	errChan <- http.ListenAndServe(":8080", handler)
    	}()

    	log.Fatalln(<-errChan)
}

