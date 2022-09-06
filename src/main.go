package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	userInput := make(chan int)
	for {
		go printMenu(userInput)
		handleInput(userInput)
	}
}

func handleInput(userInput chan int) {

	select {
	case choice := <-userInput:
		switch choice {
		case 1:
			serverOutput := make(chan string)
			go getTime(serverOutput)
			fmt.Println("Time is", <-serverOutput)
		case 2:
			sayHelloWorld()
		case 9:
			fmt.Println("Thanks for playing")
			os.Exit(0)
			return
		}
	}
}

func printMenu(inputChan chan int) {
	fmt.Println("Please choose an option")
	fmt.Println("\t1.) Print time")
	fmt.Println("\t2.) Say 'Hello World'")
	fmt.Println("\t9.) Exit")
	var input string
	fmt.Scanln(&input)
	intVal, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Could not parse input %v", err)
		inputChan <- 9
		return
	}
	inputChan <- intVal
}
