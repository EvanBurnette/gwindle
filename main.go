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

// func buildSet(phrase[]string, wordsLeft []string) []string {
// 	if len(wordsLeft) == 0 {
// 		return phrase
// 	}
// 	for _, phrase := range phrases {
// 		//filter out words with common letters

// 	}
// 	return phrases
// }

func main() {
	dat, err := os.ReadFile("fives.csv")
	check(err)
	words := strings.Split(strings.Trim(string(dat), "\n"), "\n")
	var filtFives []string
	//filter out words with duplicate letters
	//and filter out words with more than one vowel
	for _, word := range words {
		letterCount := make(map[string]int)
		add := true
		letters := strings.Split(word, "")
		sort.Strings(letters)
		vowels := "aeiou"
		numVowels := 0
		for _, letter := range letters {
			letterCount[letter] += 1
			if strings.Contains(vowels, letter) {
				numVowels += 1
				if numVowels > 1 {
					add = false
					break
				}
			}
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

	//remove anagrams
	lastWord := ""
	var noAnagrams []string
	for _, word := range filtFives {
		if lastWord != word {
			noAnagrams = append(noAnagrams, word)
		}
		lastWord = strings.Clone(word)
	}
	// fmt.Println(strings.Join(noAnagrams, "\n"))

	wordLists := make(map[string][]string)

	vowels := strings.Split("aeiou", "")

	for _, w := range noAnagrams {
		for _, v := range vowels {
			if strings.Contains(w, v) {
				wordLists[v] = append(wordLists[v], w)
				break
			}
		}
	}
	passes := len(wordLists["a"])
	fmt.Println(len(wordLists["a"]) * len(wordLists["e"]) * len(wordLists["i"]) * len(wordLists["o"]) * len(wordLists["u"]))
	// var megaList []string
	mostLetters := 24
	for j, a := range wordLists["a"] {
		fmt.Printf("%d/%d\n", j, passes)
		for _, e := range wordLists["e"] {
			if strings.ContainsAny(e, a) {
				continue
			}
			for _, i := range wordLists["i"] {
				if strings.ContainsAny(i, a+e) {
					continue
				}
				for _, o := range wordLists["o"] {
					if strings.ContainsAny(o, a+e+i) {
						continue
					}
					for _, u := range wordLists["u"] {
						phrase := a + e + i + o + u
						uniques := make(map[rune]rune)
						for _, letter := range phrase {
							uniques[letter] = letter
						}
						length := len(uniques)
						if length >= mostLetters {
							fmt.Printf("%d %s %s %s %s %s\n", length, a, e, i, o, u)
							if length > mostLetters {
								mostLetters = length
							}
						}
					}
				}
			}
		}
	}
	// mostLetters := 0
	// for _, phrase := range megaList {
	// 	uniques := make(map[rune]int)
	// 	for _, letter := range phrase {
	// 		uniques[letter] += 1
	// 	}
	// 	length := len(uniques)
	// 	if length >= mostLetters {
	// 		fmt.Println(phrase)
	// 		mostLetters = length
	// 	}
	// }
}
