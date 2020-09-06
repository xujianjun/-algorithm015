package Week_01

func rotate(nums []int, k int) {
	l := len(nums)
	diff := k % l

	copy(nums, append(nums[l-diff:], nums[:l-diff]...))
	return
}
