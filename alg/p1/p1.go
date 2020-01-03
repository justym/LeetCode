package p1

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				indices[0] = i
				indices[1] = j
				return indices
			}
		}
	}
	return indices
}
