package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello world")

	text := "I am living in Hawaii"

	words := strings.Split(text, " ")

	fmt.Println(words)

	list_of_words := make(rankWords, 0)

	for _, word := range words {
		cons, vows := countConsAndVows(word)
		fmt.Println(word, "has", cons, "consonants and", vows, "vowels")
		diff := int(math.Abs(float64(cons - vows)))
		list_of_words = append(list_of_words, wordSorter{diff, word})
	}

	sort.Sort(rankWords(list_of_words))
	fmt.Println(list_of_words)
	for _, word := range list_of_words {
		fmt.Println("Word:", word.word)
	}
}

type wordSorter struct {
	rank int
	word string
}

type rankWords []wordSorter

func (w rankWords) Len() int {
	return len(w)
}

func (w rankWords) Less(i, j int) bool {
	if w[i].rank == w[j].rank {
		fmt.Println(strings.ToLower(w[i].word), "compare to", strings.ToLower(w[j].word))
		return strings.ToLower(w[i].word) < strings.ToLower(w[j].word)
	} else {
		return w[i].rank < w[j].rank
	}
}

func (w rankWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func countConsAndVows(s string) (int, int) {
	cons := 0
	vows := 0
	for _, c := range s {
		if isVowel(c) {
			vows++
		} else {
			cons++
		}
	}
	return cons, vows
}

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
