package main

import (
	"log"
	"os/exec"
	"time"
)

func getTime(output chan string) {
	output <- time.Now().String()
}

func sayHelloWorld() {
	cmd := exec.Command("say", "Hello World")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
