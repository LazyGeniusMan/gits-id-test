package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Print("Input Bilangan: ")
	var inputS string
	fmt.Scanln(&inputS)
	inputI, err := strconv.Atoi(inputS)
	if err != nil {
		fmt.Println("Input bukan bilangan!")
		log.Fatal(err)
	}
	switch {
	case inputI%3 == 0 && inputI%5 == 0:
		fmt.Print("Hello World")
	case inputI%3 == 0:
		fmt.Print("Hello")
	case inputI%5 == 0:
		fmt.Print("World")
	default:
		fmt.Print("Tidak memenuhi kondisi manapun")
	}
}
