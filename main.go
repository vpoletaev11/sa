package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	check := cutter(ipOutputByte)
	fmt.Println(check)
}

// cut ip link output to strings
func cutter(text []byte) []string {
	clean := strings.Replace(string(text), "\n    ", " ", -1)
	clean = clean[:len(clean)-1]
	return strings.Split(clean, "\n")
}
