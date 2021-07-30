package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func reverse(input string) string {
	var result string
	runes := []rune(input)
	for i := len(runes) - 1; i != -1; i-- {
		result += string(runes[i])
	}
	return result
}

// convert to lowercase, remove special character and whitespace
func cleanString(str string) string {
	var result string
	str = strings.ReplaceAll(str, " ", "")
	reg := regexp.MustCompile("[^a-z0-9]+")
	result = reg.ReplaceAllString(strings.ToLower(str), "")
	return result
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Kalimat: ")
	input, _ := consoleReader.ReadString('\n')
	reversed := reverse(input)
	if cleanString(input) == cleanString(reversed) {
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}
}
