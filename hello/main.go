package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("hello JI-2016!\n")
	fmt.Printf("running pkg-config --version systemd now...\n")
	cmd := exec.Command("pkg-config", "--version", "systemd")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
