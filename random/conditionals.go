//go:build ignore

package main

import "fmt"

func main() {
	messageLen := 10
	maxLen := 20

	if messageLen < maxLen {
		fmt.Printf("Your message is fine! Current message length: %v\n", messageLen)
	} else if messageLen == maxLen {
		fmt.Printf("Spot on! Current message length: %v\n", messageLen)
	} else {
		fmt.Printf("Your message is too long! Current message length: %v\n", messageLen)
	}

	email := "sth@gmail.com"

	// emailLen can be only used within if statement
	if emailLen := len(email); emailLen < 20 {
		fmt.Printf("Email is valid!\n")
	}

}
