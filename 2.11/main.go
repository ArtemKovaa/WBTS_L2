package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func findAnagrams(strs []string) map[string][]string {
	tempRes := make(map[string][]string)

	for _, str := range strs {
		lower := strings.ToLower(str)
		r := []rune(lower)
		slices.Sort(r)
		key := string(r)
		tempRes[key] = append(tempRes[key], lower)
	}

	res := make(map[string][]string)
	for _, innerSlice := range tempRes {
		if len(innerSlice) > 1 {
			key := innerSlice[0]
			sort.Strings(innerSlice)
			res[key] = innerSlice
		}
	}

	return res
}

func main() {
	fmt.Println(findAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}))
}
