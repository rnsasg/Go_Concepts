package main

import (
	"fmt"
	"strconv"
)

var (
	mstr = map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "0"},
		7: []string{"p", "q", "r"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
)

func main() {
	// fmt.Println(mstr)
	// fmt.Println(letterCombinations("1"))
	//fmt.Println(combineTwo("1", "2"))
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	var res []string
	tempRes := make([][]string, len(digits))
	if len(digits) == 0 {
		return res
	}

	if len(digits) == 1 {
		r, _ := strconv.Atoi(digits)
		return mstr[r]
	}

	for i := 1; i < len(digits); i++ {
		fmt.Println(string(digits[:i]), string(digits[i]))
		tempRes[i] = combineTwo(string(digits[:i]), string(digits[i]))
		r, _ := strconv.Atoi(string(digits[:i+1]))
		mstr[r] = tempRes[i]
	}
	return tempRes[len(digits)-1]
}

func combineTwo(d1 string, d2 string) []string {
	var res []string
	t, _ := strconv.Atoi(d1)
	str1 := mstr[t]
	fmt.Println(str1)
	t, _ = strconv.Atoi(d2)
	str2 := mstr[t]
	fmt.Println(str2)
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			res = append(res, fmt.Sprintf("%s%s", str1[i], str2[j]))
		}
	}
	return res
}
