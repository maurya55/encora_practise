// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	twoSum(nums, target)

}

func twoSum(nums []int, target int) []int {

	var result = []int{}

	for i := 0; i < len(nums); i++ {

		for j := 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				result = append(result, nums[i], nums[j])
			}
		}

	}

	return result

}
