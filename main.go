package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("fives.csv")
	check(err)
	words := strings.Split(strings.Trim(string(dat), "\n"), "\n")
	var filtFives []string
	//filter out words with duplicate letters
	for _, word := range words {
		letterCount := make(map[rune]int)
		add := true
		for _, letter := range word {
			letterCount[letter] += 1
			if letterCount[letter] > 1 {
				add = false
				break
			}
		}
		if add {
			filtFives = append(filtFives, word)
		}
	}

	//anagrams
	anagrams := make(map[string][]string)
	var nograms []string
	for _, word := range filtFives {
		sorted := strings.Split(word, "")
		sort.Strings(sorted) //sort package sorts in place and has no return value
		key := strings.Join(sorted, "")
		_, hasKey := anagrams[key]
		if hasKey {
			nograms = append(nograms, word)
		}
		anagrams[key] = append(anagrams[key], word)
	}
	fmt.Println(strings.Join(nograms, "\n"))

	//create solution channel
	c := make(chan string)

	//create and test combinations for each word in list
	for i, word := range nograms {
		go combineAndTest(i, word, nograms, c)
	}

	for {
		v := <- c
		fmt.Println(v)
	}
}

func combineAndTest(i int, word string, list []string, c chan string) {
	c <- word
}