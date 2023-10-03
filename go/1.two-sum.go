package twoSum

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for firstIndex, firtsValue := range nums {
		for secondIndex, secondValue := range nums {
			if firstIndex == secondIndex {
				continue
			}

			if (firtsValue + secondValue) == target {
				return []int{firstIndex, secondIndex}
			}
		}
	}

	return []int{}
}

// @lc code=end
