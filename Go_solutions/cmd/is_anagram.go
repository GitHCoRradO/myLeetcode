package main

import "sort"

func main() {

}

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}
	sortedS := sortString(s)
	sortedT := sortString(t)
	if sortedT == sortedS {
		return true
	}
	return false
}

func sortString(s string) string {
	runeArr := []int{}
	for _, c := range s {
		runeArr = append(runeArr, int(c))
	}
	sort.Ints(runeArr)
	rune32Arr := []rune{}
	for _, c := range runeArr {
		rune32Arr = append(rune32Arr, int32(c))
	}
	return string(rune32Arr)
}

