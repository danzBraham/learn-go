package main

import (
	"fmt"
)

func main() {
	names := []string{"Zidan", "Zoe", "Baki", "Zidan"}
	fmt.Println(getNameCounts(names))
}

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}

		initial := rune(name[0])

		if _, ok := counts[initial]; !ok {
			counts[initial] = make(map[string]int)
		}

		counts[initial][name]++
	}
	return counts
}
