package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//contains parameters for adapter
type adapter struct {
	number string
	name   string
	mac    string
	mode   string
}

func main() {
	ipOutputByte, err := exec.Command("ip", "link").Output()
	if err != nil {
		log.Fatal(err)
	}
	cutted := cutter(ipOutputByte)
	adapters := aggregator(cutted)
	printAdapters(adapters)
}

// cut ip link output to strings
func cutter(text []byte) []string {
	clean := strings.Replace(string(text), "\n    ", " ", -1)
	// delete last "\n" (redundunt)
	clean = clean[:len(clean)-1]
	return strings.Split(clean, "\n")
}

// cuts out data from lines to slice of adapter structs
func aggregator(lines []string) []adapter {
	adapters := []adapter{}
	for _, value := range lines {
		// err handle
		words := wordExtractor(value)
		// output of "ip link" assumed to be hardcoded
		a := adapter{
			number: words[0],
			name:   words[1],
			mac:    words[16],
			mode:   words[10],
		}
		adapters = append(adapters, a)
	}
	return adapters
}

// create slice of words from line, remove ":" from last element of words
func wordExtractor(line string) []string {
	// create slice of words from line
	words := strings.Split(line, " ")
	// remove ":"
	for i, value := range words[:2] {
		value = value[:len(value)-1]
		words[i] = value
	}
	return words
}

func printAdapters(adapters []adapter) {
	// find adapter with longest name
	lenLongestName := 0
	for _, a := range adapters {
		if len(a.name) > lenLongestName {
			lenLongestName = len(a.name)
		}
	}
	for i := range adapters {
		// add ":" to number
		a := adapters[i]
		number := a.number + ":"
		//
		name := ""
		if len(a.name) < lenLongestName {
			countSpaces := lenLongestName - len(a.name)
			name = a.name + strings.Repeat(" ", countSpaces)
		} else {
			name = a.name
		}
		fmt.Println(number, name, a.mac, a.mode)
	}
}
