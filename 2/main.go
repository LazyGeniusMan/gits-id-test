package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Print("Input Email: ")
	var input string
	userP := "^[a-zA-Z0-9_.+]{1,20}[@].*$"                                //for email username; 1-20 characters; underscore, dot, and plus allowed
	tldP := "^.*[@][a-zA-Z0-9-]{2,63}[.](?:id|co.id)$"                    //for TLD;  only .id and .co.id allowed
	fullP := "^[a-zA-Z0-9_.+]{1,20}[@][a-zA-Z0-9-]{2,63}[.](?:id|co.id)$" //full pattern
	fmt.Scanln(&input)
	fullV, _ := regexp.MatchString(fullP, input)
	if fullV == true {
		fmt.Println("Format Email Benar")
	} else {
		fmt.Println("Format Email Salah")
		userV, _ := regexp.MatchString(userP, input)
		if userV == false {
			fmt.Println("Username maksimal 20 karakter")
		}
		tldV, _ := regexp.MatchString(tldP, input)
		if tldV == false {
			fmt.Println("TLD yang diperbolehkan hanya .id dan .co.id")
		}
	}
}
