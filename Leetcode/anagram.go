package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(anagram("nat", "tan"))
	//fmt.Println(anagram("eat", "tea"))
	//fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(getSignature("nat"))
}

func anagram(str1, str2 string) bool {

	mp := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		mp[str1[i]]++
	}

	for i := 0; i < len(str2); i++ {
		if mp[str2[i]] > 0 {
			mp[str2[i]]--
			if mp[str2[i]] == 0 {
				delete(mp, str2[i])
			}
			continue
		} else {
			return false
		}
	}

	return len(mp) == 0
}

func getSignature(str string) string {
	var charArr [26]int

	for i := 0; i < len(str); i++ {
		charArr[str[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		fmt.Print(rune(charArr[i] + 'a'))
	}

	fmt.Println("Returns")

	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if charArr[i] != 0 {
			for j := charArr[i]; j > 0; j-- {
				fmt.Print(charArr[i] + 'a')
				sb.WriteString(string(charArr[i] + 'a'))
			}
		}
	}
	return sb.String()
}

func groupAnagrams(strs []string) [][]string {
	processed := make(map[int]bool)
	res := [][]string{}
	temp := []string{}

	for i := 0; i < len(strs); i++ {
		if processed[i] == true {
			continue
		}
		temp := append(temp, strs[i])
		for j := i + 1; j < len(strs); j++ {
			if processed[j] == false && anagram(strs[i], strs[j]) == true {
				temp = append(temp, strs[j])
				processed[j] = true
			}
		}
		res = append(res, temp)
	}

	return res
}
