package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("catsandog", []string{"cat", "sand", "cats", "dog"}))
}

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, word := range wordDict {
		set[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && set[s[j:i-1]] {
				dp[i-1] = true
			}
		}
	}
	return dp[len(s)]
}
