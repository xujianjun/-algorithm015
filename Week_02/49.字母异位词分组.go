package Week_02

import "sort"

type bytes []byte

// 自定义排序规则
func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}

	record := map[string]int{}
	for _, v := range strs {
		vbyte := bytes(v)
		sort.Sort(vbyte)
		vsort := string(vbyte)
		if index, ok := record[vsort]; ok {
			ret[index] = append(ret[index], v)
		} else {
			record[vsort] = len(ret)
			ret = append(ret, []string{v})
		}
	}
	return ret
}
