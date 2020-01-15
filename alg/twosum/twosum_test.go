package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	num := make([]int, 4, 4)
	num = []int{2, 7, 11, 15}
	target := 9
	ans := make([]int, 2, 2)
	ans = []int{0, 1}
	com := TwoSum(num, target)

	for i := 0; i < len(ans); i++ {
		if com[i] != ans[i] {
			t.Logf("com[%d] cannot pass", i)
			t.Error(t)
		}
	}
}
