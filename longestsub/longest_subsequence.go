package longestsub

import (
	"sort"
	"strings"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

// IsSubsequence determines whether word is a subsequence of ref.
func IsSubsequence(ref string, word string) bool {
	letters := strings.Split(word, "")
	for _, letter := range letters {
		in := strings.Index(ref, letter)
		if in >= 0 {
			ref = string(ref[in+1:])
		} else {
			return false
		}
	}
	return true
}

// FindLongestSubsequence determines the longest subsequence of ref from words.
func FindLongestSubsequence(ref string, words []string) string {
	longest := ""
	sort.Sort(byLength(words))
	for _, word := range words {
		longestLength := len(longest)
		if len(word) > longestLength && IsSubsequence(ref, word) == true {
			longest = word
		}
	}
	return longest
}
