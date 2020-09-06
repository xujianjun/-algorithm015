package Week_02

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm := map[string]int{}
	for i := 0; i < len(s); i++ {
		sv := string(s[i])
		if _, ok := sm[sv]; ok {
			sm[sv]++
		} else {
			sm[sv] = 1
		}
	}
	fmt.Println(sm)

	for j := 0; j < len(t); j++ {
		tv := string(t[j])
		if _, ok1 := sm[tv]; ok1 {
			sm[tv]--
		} else {
			sm[tv] = 1
		}
	}
	fmt.Println(sm)

	res := true
	for _, vv := range sm {
		if vv != 0 {
			res = false
			break
		}
	}
	return res
}
