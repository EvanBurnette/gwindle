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
		letterCount := make(map[string]int)
		add := true
		letters := strings.Split(word, "")
		sort.Strings(letters)
		for _, letter := range letters {
			letterCount[letter] += 1
			if letterCount[letter] > 1 {
				add = false
				break
			}
		}
		if add {
			filtFives = append(filtFives, strings.Join(letters, ""))
		}
	}
	sort.Strings(filtFives)
	fmt.Println(strings.Join(filtFives, "\n"))
	//remove anagrams
}
