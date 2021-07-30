package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findHour(input string) int {
	hourP := "^(1[0-2]|0?[1-9])"
	reg := regexp.MustCompile(hourP)
	hourS := reg.FindString(input)
	hourI, _ := strconv.Atoi(hourS)
	return hourI
}

func findAmPm(input string) string {
	ampmP := "(?:AM|PM)"
	reg := regexp.MustCompile(ampmP)
	ampm := reg.FindString(input)
	return ampm
}

//02-09 ??
func convert(tInput string, hour int, ampm string) string {
	var hourConverted int
	if ampm == "PM" {
		hourConverted = hour + 12
	} else {
		hourConverted = hour
	}
	tAmPmConverted := strings.ReplaceAll(tInput, ampm, "")
	tConverted := strings.ReplaceAll(tAmPmConverted, strconv.Itoa(hour), strconv.Itoa(hourConverted))
	return tConverted
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Contoh Format 10:11:12 PM")
	fmt.Print("Input Waktu: ")
	tInput, _ := consoleReader.ReadString('\n')
	//tP := "^(1[0-2]|0?[1-9]):[0-5][0-9]:[0-5][0-9]\\s(?:AM|PM)$"
	hour := findHour(tInput)
	ampm := findAmPm(tInput)
	tConverted := convert(tInput, hour, ampm)
	fmt.Print(tConverted)
}
