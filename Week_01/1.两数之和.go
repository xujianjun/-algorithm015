package Week_01

func twoSum(nums []int, target int) []int {
	res := []int{}

	hashTar := map[int]int{}
	for i, v := range nums {
		if j, ok := hashTar[target-v]; ok {
			res = append(res, j, i)
			break
		}
		hashTar[v] = i
	}

	return res
}
