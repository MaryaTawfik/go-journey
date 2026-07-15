package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func removePunctuation(s string) string {
	clean := ""

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			clean += string(r)
		}
	}

	return clean
}

func main() {
	fmt.Println("Enter the text:")

	var s string
	// fmt.Scanln(&s)
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.ToLower(s)

	newText := removePunctuation(s)

	words := strings.Split(newText, " ")

	wordCounts := map[string]int{}

	for _, word := range words {
		if word == "" {
			continue
		}

		_, exists := wordCounts[word]

		if exists {
			wordCounts[word]++
		} else {
			wordCounts[word] = 1
		}
	}

	fmt.Println(wordCounts)
}