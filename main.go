package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bbalet/stopwords"
)

var words = "Please create a small service that accepts as input a body of text, such as that from a book, and returns the top ten most-used words along with how many times they occur in the text."

func topten(words string) []string {
	input := strings.Fields(words)
	count := make(map[string]int)

	for _, word := range input {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		word := reg.ReplaceAllString(word, "")
		_, match := count[word]
		if match {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}

	keys := make([]string, 0, len(count))
	out := make([]string, 0)
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for i, name := range keys {
		var s = strconv.Itoa(count[name])
		if i != 10 {
			out = append(out, name+":"+s)
		} else {
			break
		}
	}
	return out
}

// top ten using remove stop words, i'm using https://github.com/bbalet/stopwords for list of stop words in english
// this function is alternative if test case needs to remove stop words like (a, an, is ,are, the, then, etc...)
func toptenusestopwords(words string) []string {
	cleanContent := stopwords.CleanString(words, "en", true)

	input := strings.Fields(cleanContent)
	count := make(map[string]int)

	for _, word := range input {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		word := reg.ReplaceAllString(word, "")
		_, match := count[word]
		if match {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}

	keys := make([]string, 0, len(count))
	out := make([]string, 0)
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for i, name := range keys {
		var s = strconv.Itoa(count[name])
		if i != 10 {
			out = append(out, name+":"+s)
		} else {
			break
		}
	}
	return out
}

func main() {
	out := topten(words)
	fmt.Println(out)
	outWithRemoval := toptenusestopwords(words)
	fmt.Println(outWithRemoval)
}
