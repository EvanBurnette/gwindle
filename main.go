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
	dat, err := os.ReadFile("test.csv")
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
	c := make(chan string)

	//create waitgroup
	var wg sync.WaitGroup

	//create and test combinations for each word in list
	//TODO fix the repetition here (reduce to recursive calls?)
	//This section causes so many off by one errors and is hard to refactor
	for i, iword := range nograms[:len(nograms)-4] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, jword := range nograms[i+1 : len(nograms)-3] {
				if strings.ContainsAny(iword, jword) {
					continue
				}
				for _, kword := range nograms[i+2 : len(nograms)-2] {
					if strings.ContainsAny(iword+jword, kword) {
						continue
					}
					for _, lword := range nograms[i+3 : len(nograms)-1] {
						if strings.ContainsAny(iword+jword+kword, lword) {
							continue
						}
						for _, mword := range nograms[i+4:] {
							if strings.ContainsAny(iword+jword+kword+lword, mword) {
								continue
							}
							c <- iword + " " + jword + " " + kword + " " + lword + " " + mword
						}
					}
				}
			}
		}()
	}
	go func() {
		for {
			v := <-c
			fmt.Println(v)
		}
	}()
	wg.Wait()
	time.Sleep(time.Millisecond * 100)
	close(c)
}
