package main

import (
	"fmt"
	"strings"
	"unicode"
)

var words = "Go! is fun and go? is Fast"

type wordFrequencyMap map[string]int

func main() {
	text := "Go! 1.23 #awesome"
	testCleanStrings := removeAllPunctuation(text)
	fmt.Println(testCleanStrings)

	cleanStrings := removeAllPunctuation(words)
	wordsSlice := convertWordsToSlice(cleanStrings)
	wordFrequency := countWords(wordsSlice)

	fmt.Println("Word Frequency")
	fmt.Println("---------------")
	for key, value := range wordFrequency {
		fmt.Printf("%s : %d\n", key, value)
	}
	fmt.Println("---------------")
	mostFrequent, count := findMostFrequentWord(wordFrequency)
	fmt.Println("Most frequent word:")
	fmt.Printf("%s (%d)\n", mostFrequent, count)
	fmt.Println("---------------")
	totalWords, uniqueWords := countTotalNumberOfWords(wordFrequency)
	fmt.Printf("Total words: %d\n", totalWords)
	fmt.Printf("Unique words: %d\n", uniqueWords)

}

func wordsToLowerCase(words string) string {
	return strings.ToLower(words)
}
func convertWordsToSlice(words string) []string {
	return strings.Fields(wordsToLowerCase(words))
}

func countWords(words []string) wordFrequencyMap {
	wordFrequency := wordFrequencyMap{}
	for _, word := range words {
		wordFrequency[word] = wordFrequency[word] + 1
	}
	return wordFrequency

}

func findMostFrequentWord(wordsMap wordFrequencyMap) (string, int) {
	mostFrequent := ""
	highestFrequency := 0
	for word, frequency := range wordsMap {
		if frequency > highestFrequency {
			mostFrequent = word
			highestFrequency = frequency

		}
	}
	return mostFrequent, highestFrequency

}

func countTotalNumberOfWords(wordsMap wordFrequencyMap) (int, int) {
	totalWords := 0
	uniqueWords := 0
	for _, frequency := range wordsMap {
		uniqueWords++
		totalWords += frequency

	}
	return totalWords, uniqueWords

}

func removeAllPunctuation(words string) string {
	cleanStrings := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, words)

	return cleanStrings

}
