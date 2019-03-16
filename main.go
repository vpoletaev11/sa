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
	mode   string
}

func main() {
	ipOutputByte, err := exec.Command("ip", "link", "show").Output()
	if err != nil {
		log.Fatal(err)
	}
	cutted := cutter(ipOutputByte)
	fmt.Println(wordExtractor(cutted[0], 0))
}

// cut ip link output to strings
func cutter(text []byte) []string {
	clean := strings.Replace(string(text), "\n    ", " ", -1)
	clean = clean[:len(clean)-1]
	return strings.Split(clean, "\n")
}

// func agregator(value []string) []adapter {
// }

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
