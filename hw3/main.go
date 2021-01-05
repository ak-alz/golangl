package main

import (
	"fmt"
	"sort"
)

func main() {
	text := "Очень длинный текст о том, как хорошо быть котиком или кроликом, но котиком значительно лучше, котики они такие замечательные дители длинного текста, ведь текст о том что котики они такие"

	fmt.Println(getCount(text, 10))
}

func getCount(text string, count int) PairList {

	frequencyKeys := map[string]int{}

	frequencyKeys["вот"] = 1
	frequencyKeys["тут"] = 12

	orderedList := rankByWordCount(frequencyKeys)

	if orderedList.Len() > count {
		orderedList = orderedList[0:count]
	}

	return orderedList
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

// Pair is..
type Pair struct {
	Key   string
	Value int
}

//PairList is
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
