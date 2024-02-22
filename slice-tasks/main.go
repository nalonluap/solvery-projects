package main

import (
	"fmt"
	"sort"
)

// Функция для сортировки строки
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// Функция для построения массива анаграмм
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		fmt.Println(str, sortedStr)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func main() {
	s := []string{"abc", "bac", "defg", "a", "fegd", "degf"}
	anagramGroups := groupAnagrams(s)
	fmt.Println(anagramGroups)
}
