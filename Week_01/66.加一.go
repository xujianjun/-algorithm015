package Week_01

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		target := digits[i] + 1
		if target < 10 {
			digits[i] = target
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
