package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Aliquip magna irure eiusmod sit elit ipsum incididunt qui labore fugiat culpa ut commodo irure Excepteur duis Lorem dolor adipisicing dolor officia mollit voluptate Consectetur exercitation occaecat laboris voluptate amet aliqua esse ex anim enim Elit exercitation officia tempor consectetur consequat reprehenderit dolor labore ex veniam minim fugiat consequat Aliquip velit voluptate qui aute enim Magna labore sint nulla magna ipsum"
	words := strings.Split(str, " ")
	var wordsCountBySize map[int]int = map[int]int{}
	getWordsCountBySize(words, &wordsCountBySize)
	maxCount, maxSize := getMaxCountBySize(wordsCountBySize)
	fmt.Printf("The max occurence of words by size of the word = %d (word size = %d)\n", maxCount, maxSize)
}

func getWordsCountBySize(words []string, wordsCountBySize *map[int]int) {
	for _, word := range words {
		(*wordsCountBySize)[len(word)]++
	}
}

func getMaxCountBySize(wordsBySize map[int]int) (int, int) {
	maxCount := 0
	maxSize := 0
	for size, count := range wordsBySize {
		if count > maxCount {
			maxCount = count
			maxSize = size
		}
	}
	return maxCount, maxSize
}
