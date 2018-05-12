package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"	
	"sort"
)

func SortUserIp(sortIp string) string {
	s := strings.Split(sortIp, "")

	sort.Strings(s)

	return strings.Join(s, "")
}

func main() {
	fmt.Printf("Hello world\n")
	
	scanner := bufio.NewScanner(os.Stdin)
	
	var text string
	fmt.Print("Enter your text ")
	scanner.Scan()
	text = scanner.Text()

	fmt.Printf("Your enterest text was %s\n", text)
	
	SortedText := SortUserIp(text)
	fmt.Printf("Sorted text is %s\n", SortedText)
}
