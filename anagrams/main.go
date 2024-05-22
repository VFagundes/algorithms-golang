package main

import (
	"fmt"
	"sort"
	"strings"
)

//#Anagrams:
//#  adapt,aptad,dream,armed
// #output [[adapt,aptad],[dream,armed]]

func sortString(w string) string {
	result := strings.Split(w, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

func allAnagrams(words []string) [][]string {
	an := make(map[string][]string)

	for _, w := range words {
		sorted := sortString(w)
		an[sorted] = append(an[sorted], w)
	}

	result := make([][]string, 0)

	for _, group := range an {
		result = append(result, group)
	}
	return result

}
func anagramRefactor(words []string) [][]string {
	result := make([][]string, 0)
	wordM := make(map[string][]string)
	for _, w := range words {
		sorted := []rune(w)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		wordM[string(sorted)] = append(wordM[string(sorted)], w)

	}
	for _, v := range wordM {
		result = append(result, v)
	}

	return result

}
func main() {
	input := strings.Split("adapt,aptad,dream,armed,tralalala", ",")
	var a, b [][]string
	a = anagramRefactor(input)
	b = allAnagrams(input)
	fmt.Println(a)
	fmt.Println(b)

}
