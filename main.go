package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//contain parameters for adapter
type adapter struct {
	number int
	name   string
	mac    string
	mode   string
}

func main() {
	ipOutputByte, err := exec.Command("ip", "link", "show").Output()
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
	clean = clean[:len(clean)-1]
	return strings.Split(clean, "\n")
}

// cuts out data from value to slice of adapter structs
func aggregator(value []string) []adapter {
	adapters := []adapter{}
	for i := range value {
		number, _ := strconv.Atoi(wordExtractor(value[i], 0))
		a := adapter{
			number: number,
			name:   wordExtractor(value[i], 1),
			mac:    wordExtractor(value[i], 16),
			mode:   wordExtractor(value[i], 10),
		}
		adapters = append(adapters, a)
	}
	return adapters
}

// extract word by position
func wordExtractor(value string, position int) string {
	// create slice of spaces for value
	spaces := []int{}
	for i := range value {
		if string(value[i]) == " " {
			spaces = append(spaces, i)
		}
	}
	// create slice of words for value
	words := []string{}
	prevSpace := 0
	for _, index := range spaces {
		word := ""
		if position == 0 {
			word = value[:index]
		} else {
			word = value[prevSpace+1 : index]
		}
		if string(word[len(word)-1]) == ":" {
			words = append(words, word[:len(word)-1])
		} else {
			words = append(words, word)
		}
		prevSpace = index
	}
	return words[position]
}

func printAdapters(adapters []adapter) {
	// find adapter with longest name
	lenLongestName := 0
	for i := range adapters {
		a := adapters[i]
		if len(a.name) > lenLongestName {
			lenLongestName = len(a.name)
		}
	}
	for i := range adapters {
		// add ":" to number
		a := adapters[i]
		number := strconv.Itoa(a.number) + ":"
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
