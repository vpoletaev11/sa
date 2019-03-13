package main

import (
	"fmt"
	"log"
	"os/exec"
)

//contain parameters for adapter
type adapter struct {
	number int
	name   string
	mac    string
}

func main() {
	ipOutputByte, err := exec.Command("ip", "link", "show").Output()
	if err != nil {
		log.Fatal(err)
	}
	check := prettyOut(ipOutputByte)
	spaces := adapterSrtuct(check)
	fmt.Println(check, spaces)
}

//return one string for one adapter from ip link
func prettyOut(out []byte) []string {
	//massive with strings for every adapter
	var strings []string
	//counter of positions '\n'
	lastN := 0
	//cutting output from ip link to strings
	for i := 0; i < len(out); i++ {
		if out[i] == '\n' {
			strings = append(strings, string(out[lastN:i]))
			lastN = i
		}
	}
	//join strings for every single adapter
	var adapterStr []string
	for i := 0; i < len(strings); i += 2 {
		if i == 0 {
			i++
		}
		adapterStr = append(adapterStr, strings[i-1]+strings[i])
	}
	return adapterStr
}

//make slice of adapters from slice of strings
func adapterSrtuct(s []string) [][]int {
	var spaces [][]int
	for i := 0; i < len(s); i++ {
		currentStr := string(s[i])
		var currentStrSpaces []int
		for a := 0; a < len(currentStr); a++ {
			if string(currentStr[a]) == " " {
				currentStrSpaces = append(currentStrSpaces, a+1)
			}
			if a == len(currentStr)-1 {
				spaces = append(spaces, currentStrSpaces)
			}
		}
	}
	return spaces
}
