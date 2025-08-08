package leetcode

func findAnagrams(s string, p string) []int {
    var result []int
    if len(s) < len(p) { return result }
    var pCount [26]int
    var sCount [26]int
    for i := 0; i < len(p); i++ {
        pCount[p[i] - 'a']++
        sCount[s[i] - 'a']++
    }

    if pCount == sCount { 
		result = append(result, 0) 
	}

    for i := len(p); i < len(s); i++ {
        sCount[s[i] - 'a']++
        sCount[s[i - len(p)] - 'a']--
        if pCount == sCount { result = append(result, i - len(p) + 1) }
    }
	return result
}