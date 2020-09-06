package Week_01

func removeDuplicates(nums []int) int {
	i := 0
	j := 0
	for ; i < len(nums); i++ {
		if i == 0 {
			nums[j] = nums[i]
			j++
			continue
		}

		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	// 处理后续的
	for k := j; k < i; k++ {
		nums[k] = 0
	}

	return j
}
