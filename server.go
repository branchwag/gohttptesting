package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the website. The whole thing. Really.\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Haiiiii\n")
}

func main(){
	//create logs
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "LOG: ", log.LstdFlags|log.Lshortfile)
	logger.Println("Logger initiated.")

	//server
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	httperr := http.ListenAndServe(":3333", nil)

	if errors.Is(httperr, http.ErrServerClosed){
		logger.Printf("server closed\n")
	} else if httperr != nil {
		logger.Printf("error starting server %s\n", httperr)
		os.Exit(1)
	}
}