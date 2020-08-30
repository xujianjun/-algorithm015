package Week_01

func moveZeroes(nums []int) {
	len := len(nums)
	i := 0
	for {
		if i >= len {
			break
		}
		if nums[i] == 0 {
			len = len - 1
			nums = append(nums[0:i], nums[i+1:]...)
			nums = append(nums, 0)
		} else {
			i++
		}
	}
	return
}
