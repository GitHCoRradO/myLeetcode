package main

func firstUniqChar(s string) int {
	charMap := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		_, exists := charMap[s[i]]
		if exists {
			continue
		} else {
			charMap[s[i]] = true
		}
		foundFlag := false
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				foundFlag = true
				break
			}
		}
		if foundFlag == false {
			return i
		}
	}
	return -1
}