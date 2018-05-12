package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Hello, world\n")
	
	scanner := bufio.NewScanner(os.Stdin)

	var text string
	
	fmt.Printf("Enter your text ");
	scanner.Scan()
	text = scanner.Text()

	fmt.Printf("Your entered text was: %s\n", text)
}
