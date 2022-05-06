package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	dat, err := os.ReadFile("fives.csv")
	if err != nil {
		panic(err)
	}
	words_raw := strings.Split(strings.Trim(strings.ReplaceAll(string(dat), "\r", ""), "\n"), "\n")
	fmt.Println(len(words_raw))

	var noDoubles []string

	func() {
		for _, word := range words_raw {
			ltCt := make(map[string]bool)
			add := true
			letters := strings.Split(word, "")
			for _, lt := range letters {
				if ltCt[lt] {
					add = false
					break
				} else {
					ltCt[lt] = true
				}
			}
			if add {
				noDoubles = append(noDoubles, word)
			}
		}
	}()
	fmt.Println(len(noDoubles))
	var words []string

	func() {
		grams := make(map[string]bool)
		for _, word := range noDoubles {
			sorted := strings.Split(word, "")
			sort.Strings(sorted)
			joined := strings.Join(sorted, "")
			if !grams[joined] {
				grams[joined] = true
				words = append(words, word)
			}
		}
	}()
	fmt.Println(len(words))

	wordGraph := make(map[string]map[string]bool)

	lenWords := len(words)
	for i, word := range words[:lenWords] {
		wordGraph[word] = make(map[string]bool)
		for _, jword := range words[i:] {
			if !strings.ContainsAny(word, jword) {
				wordGraph[word][jword] = true
			}
		}
	}

	// fmt.Println(wordGraph["yurta"])

	var keys []string
	for key, _ := range wordGraph {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, key := range keys {
		phrase := []string{key}
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			findPhrase(5, wordGraph, ch, phrase)
			// fmt.Println("finished", key, "loop")
		}(key)
	}
	go func() {
		for {
			v := <-ch
			fmt.Println(v)
		}
	}()
	wg.Wait()
	time.Sleep(time.Millisecond * 50) //Adding time for sync is hacky?
	close(ch)
}

func findPhrase(wantLen int, wordGraph map[string]map[string]bool, ch chan string, phrase []string) {
	if len(phrase) == wantLen { //base case
		ch <- strings.Join(phrase, " ")
		return
	}
	if len(phrase) == 1 { //create all combos from phrase[0]'s list
		for word, _ := range wordGraph[phrase[0]] {
			findPhrase(wantLen, wordGraph, ch, append(phrase, word))
		}
	} else {
		last := phrase[len(phrase)-1]
		for word, _ := range wordGraph[last] {
			add := true
			//for each word in the last word's word list
			for i := 0; i < len(phrase)-1; i++ {
				//if this word in the all ith words' word lists? (except for the second to last word)
				////append this word to the phrase and send it back through the function
				if !wordGraph[phrase[i]][word] {
					add = false
					break
				}
			}
			if add {
				findPhrase(wantLen, wordGraph, ch, append(phrase, word))
			}
		}
	}
}
