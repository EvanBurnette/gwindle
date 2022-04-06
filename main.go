package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
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

	//remove anagrams
	var nograms []string
	func() {
		anagrams := make(map[string]bool)
		for _, word := range filtFives {
			sorted := strings.Split(word, "")
			sort.Strings(sorted) //package "sort" sorts in place and has no return value
			key := strings.Join(sorted, "")
			if anagrams[key] != true {
				nograms = append(nograms, word)
				anagrams[key] = true
			}
		}
	}()
	//sanity check for test set that all values are still present
	fmt.Println(len(words))
	fmt.Println(len(nograms))

	//create solution channel
	ch := make(chan string)

	//create waitgroup
	var wg sync.WaitGroup

	//create and test combinations for almost every word in list
	lpEnd := len(nograms) - 3
	for i, word := range nograms[:lpEnd] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			combineTest(i, word, lpEnd, nograms, ch)
		}()
	}

	go func() {
		for {
			v := <-ch
			fmt.Println(v)
		}
	}()
	wg.Wait()
	time.Sleep(time.Millisecond * 1500) //Adding time for sync is hacky?
	close(ch)
}

func combineTest(i int, phrase string, lpEnd int, nograms []string, ch chan string) {
	//for each word in list[i+1:listLen-5-phraselen]
	j := i + 1
	for _, word := range nograms[j : lpEnd+1] {
		//if word contains any letters of phrase
		if strings.ContainsAny(phrase, word) {
			continue
		}
		if lpEnd >= len(nograms) {
			ch <- phrase + " " + word
			//add word to phrase and send to channel
		} else {
			combineTest(j, phrase+" "+word, lpEnd+1, nograms, ch)
			//add word to phrase and send it to combine and test
		}
	}
}
