package main

func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int)
	var max int
	for i, j := 0, 0; j < len(s); j++ {
		if k, ok := m[s[j]]; ok && k >= i { // if the character is repeated, like "abba"
			i = k + 1 // move i to the next of the repeated character
		}
		m[s[j]] = j
		if j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
