package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var portVar string
var hostVar string
var dirVar string

func init() {
	const (
		defaultPort  = "8080"
		portHelpText = "the port number"
		defaultHost  = "localhost"
		hostHelpText = "the host"
		defaultDir   = "."
		dirHelpText  = "the static directory"
	)
	flag.StringVar(&portVar, "port", defaultPort, portHelpText)
	flag.StringVar(&dirVar, "dir", defaultDir, dirHelpText)
	flag.StringVar(&hostVar, "host", defaultHost, hostHelpText)
}

func main() {
	flag.Parse()
	host := fmt.Sprintf("%s:%s", hostVar, portVar)
	fmt.Printf("\nListening on %s\n", host)
	err := http.ListenAndServe(host, http.FileServer(http.Dir(dirVar)))

	if err != nil {
		log.Fatal(err)
	}

}
