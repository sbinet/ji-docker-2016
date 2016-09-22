package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
)

func main() {
	http.HandleFunc("/", rootHandle)
	port := ":8080"
	log.Printf("listening on: http://localhost%s\n", port)
	err := http.ListenAndServe("0.0.0.0"+port, nil)
	if err != nil {
		log.Fatalf("error closing web server: %v\n", err)
	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello JI-2016!\n")
	fmt.Fprintf(w, "\n\n--- running external command...\n\n>>> pkg-config --version systemd\n")
	cmd := exec.Command("pkg-config", "--version", "systemd")
	cmd.Stdout = w
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
	}
}
