package palindrome

import "fmt"

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var converted string = fmt.Sprint(x)
	var reversed string
	for _, v := range converted {
		reversed = string(v) + reversed
	}

	return reversed == converted
}

// @lc code=end
