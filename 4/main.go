package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Kalimat: ")
	input, _ := consoleReader.ReadString('\n')
	var reversed string
	runes := []rune(input)
	for i := len(runes) - 1; i != -1; i-- {
		reversed += string(runes[i])
	}
	fmt.Print(reversed)
}
