package main

// This is my solution for : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// Its not optimal but i think that it's very neat (used bitswise operations and stuff)

const (
	SCHEMA    = 64
	ALPHA_NUM = 26
	NUM_NUM   = 10
)

func getPos(c byte) uint {
	if c >= 'a' && c <= 'z' {
		return uint(c) - 'a'
	} else if c >= 'A' && c <= 'Z' {
		return uint(c) - 'A' + ALPHA_NUM
	} else if c >= '0' && c <= '9' {
		return uint(c) - '0' + ALPHA_NUM + ALPHA_NUM
	} else {
		return ALPHA_NUM*2 + NUM_NUM
	}
}

func Max(f, s int) int {
	if f > s {
		return f
	}
	return s
}

func lengthOfLongestSubstring(str string) int {
	if len(str) == 0 {
		return 0
	}

	var (
		bit     uint64
		maxLen  int          = 0
		past    uint64       = 0
		temp    int          = 0
		charMap map[byte]int = make(map[byte]int, 256)
	)

	for i := 0; i < len(str); i++ {
		c := str[i]
		bit = 1 << getPos(c)

		// magik
		if (past & bit) > 0 {
			past = 0
			temp = Max(maxLen, temp)
			maxLen = i - charMap[c]
			for j := charMap[c] + 1; j <= i; j++ {
				past |= (1 << getPos(str[j]))
			}
		} else {
			past |= bit
			maxLen++
		}
		charMap[c] = i
	}

	return Max(maxLen, temp)
}
