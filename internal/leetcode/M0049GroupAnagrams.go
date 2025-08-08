package leetcode

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		key := strings.Join(s, "")
		m[key] = append(m[key], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res

}
