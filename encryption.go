package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// channels for communication
	stringChan := make(chan string)
	tower1Chan := make(chan string)
	tower2Chan := make(chan string)

	var offset int32 = 3

	// go routines for tower1 and tower2
	go tower1(stringChan, tower1Chan, offset)
	go tower2(stringChan, tower2Chan, offset)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-tower1Chan:
			fmt.Printf("Control Tower: Message from Tower 1 - %v", msg)
		case msg := <-tower2Chan:
			fmt.Printf("Control Tower: Message from Tower 2 - %v", msg)
		}
	}
}

func tower1(s chan string, t1 chan string, offset int32) {
	inputStream := bufio.NewReader(os.Stdin)
	fmt.Println("Tower 1: Enter your message for Tower 2.")
	userInput, _ := inputStream.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	fmt.Printf("Tower 1: Original String: %s", userInput)

	var secretString string

	for _, c := range userInput {
		secretString += string(c + offset)
	}

	fmt.Printf("Tower 1: Encrypted String: %s\n", secretString)

	s <- secretString
	t1 <- "Message sent to Tower 2\n"
}

func tower2(s chan string, t2 chan string, offset int32) {
	secretString := <-s
	var originalString string

	for _, c := range secretString {
		originalString += string(c - offset)
	}

	fmt.Printf("Tower 2: Decrypted Message: %s", originalString)
	t2 <- "Message received from Tower 1\n"
}
