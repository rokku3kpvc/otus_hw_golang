package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordFreq struct {
	word string
	freq int
}

var wordRe = regexp.MustCompile(`(?m)(?:\p{L}-?)+`)

func Top10(text string) []string {
	topWords := make([]string, 0, 10)
	if len(text) == 0 {
		return topWords
	}

	wordsMap := make(map[string]int)
	for _, w := range wordRe.FindAllString(text, -1) {
		w = strings.ToLower(w)
		wordsMap[w]++
	}

	wordsFreqSlice := make([]wordFreq, 0, len(wordsMap))
	for k, v := range wordsMap {
		wordsFreqSlice = append(wordsFreqSlice, wordFreq{k, v})
	}

	sort.Slice(wordsFreqSlice, func(i, j int) bool {
		if wordsFreqSlice[i].freq == wordsFreqSlice[j].freq {
			s := []string{wordsFreqSlice[i].word, wordsFreqSlice[j].word}
			return sort.StringsAreSorted(s)
		}

		return wordsFreqSlice[i].freq > wordsFreqSlice[j].freq
	})

	if len(wordsFreqSlice) > 0 {
		var l int
		if len(wordsFreqSlice) < 10 {
			l = len(wordsFreqSlice)
		} else {
			l = 10
		}

		for i := 0; i < l; i++ {
			topWords = append(topWords, wordsFreqSlice[i].word)
		}
	}

	return topWords
}
