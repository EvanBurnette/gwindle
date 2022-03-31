package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	// "reflect"
	"strconv"
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
	fmt.Println(len(words))
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
	anagrams := make(map[string]bool)
	var nograms []string
	for _, word := range filtFives {
		sorted := strings.Split(word, "")
		sort.Strings(sorted) //package "sort" sorts in place and has no return value
		key := strings.Join(sorted, "")
		if anagrams[key] != true {
			nograms = append(nograms, word)
			anagrams[key] = true;
		}
	}
	fmt.Println(len(nograms))

	//create solution channel
	c := make(chan string)

	//create and test combinations for each word in list
	for i, word := range nograms[:len(nograms)-4] {
		go combineAndTest(i, word, nograms, c)
	}

	for {
		v := <- c
		fmt.Println(v)
	}
}

func combineAndTest(i int, iword string, list []string, c chan string) {
	for _, jword := range list[i+1:len(list)-3] {
		if strings.ContainsAny(iword, jword) {
			continue
		} 
		for _, kword := range list[i+2:len(list)-2] {
			if strings.ContainsAny(iword + jword, kword) {
				continue
			}
			for _, lword := range list[i+3:len(list)-1] {
				if strings.ContainsAny(iword + jword + kword, lword) {
					continue
				}
				for _, mword := range list[i+4:] {
					if strings.ContainsAny(iword + jword + kword + lword, mword) {
						continue
					}
					c <- iword + " " + jword + " " + kword + " " + lword + " " + mword
				}
			}	
		}
	}
	c <- "finished " + strconv.Itoa(i) //TODO change this to a count channel to tell the main channel to close when all goroutines are done
}