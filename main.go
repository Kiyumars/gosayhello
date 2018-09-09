package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name!")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s!", text)
}
