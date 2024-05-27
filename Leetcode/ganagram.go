package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//fmt.Println(hash("eat"))
}

func groupAnagrams(strs []string) [][]string {
	strmp := make(map[string][]string)
	res := [][]string{}

	var hash func(string) string
	hash = func(s string) string {
		res := make([]byte, 26)
		for _, c := range s {
			res[c-'a'] += 1
		}
		return string(res)
	}

	for i := 0; i < len(strs); i++ {
		temp := hash(strs[i])
		strmp[temp] = append(strmp[temp], strs[i])
	}

	for _, v := range strmp {
		res = append(res, v)
	}
	return res
}
