package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	var network string
	var addr string

	if os.Getenv("SOCK_TYPE") == "unix" {
		const sockFile = "/tmp/httpd.sock"
		network, addr = "unix", sockFile
		os.Remove(sockFile)
	} else {
		network, addr = "tcp", ":8080"
	}
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer lis.Close()

	http.HandleFunc("/", handler)

	log.Printf("server starting on %s", addr)
	if err := http.Serve(lis, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
