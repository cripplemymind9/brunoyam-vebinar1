package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	text := scanner.Text()
	wordCount := countWords(text)
	for word, count := range wordCount {
		fmt.Printf("%s - %d\n", word, count)
	}
}

func countWords(text string) map[string]int {
	re := regexp.MustCompile(`[^\p{L}\p{N}\s]+`)
	cleanText := re.ReplaceAllString(text, "")
	cleanText = strings.ToLower(cleanText)
	words := strings.Fields(cleanText)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}