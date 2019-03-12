package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
	fmt.Println(check)
}

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
	//split strings for every single adapter
	var adapterStr []string
	for i := 0; i < len(strings); i += 2 {
		if i == 0 {
			i++
		}
		adapterStr = append(adapterStr, strings[i-1]+strings[i])
	}
	return adapterStr
}
