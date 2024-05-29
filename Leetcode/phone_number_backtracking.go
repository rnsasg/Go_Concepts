package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var (
	mstr = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
)

func letterCombinations(digits string) []string {
	ans := []string{}
	res := []byte{}

	if len(digits) == 0 {
		return ans
	}

	backtrack(digits, 0, res, &ans)

	return ans
}

func backtrack(digits string, idx int, res []byte, ans *[]string) {

	if len(res) == len(digits) {
		t := make([]byte, len(res))
		copy(t, res)
		*ans = append(*ans, string(t))
		fmt.Println(string(res))
		return
	}
	if idx >= len(digits) {
		return
	}

	//fmt.Println(int(digits[idx] - '0'))
	str, _ := mstr[int(digits[idx]-'0')]
	for i := 0; i < len(str); i++ {
		res = append(res, str[i])
		backtrack(digits, idx+1, res, ans)
		res = res[:len(res)-1]
	}
}

//"2"
//backtrack("2",0,[],[][]) --> str = "abc" --> 0 to 2
//  i = 0 --> res = {"a"}
//  backtrack("2",1,{a},[]{"a",}) --> res{}
//  i = 1 --> res = {"a"}
//  backtrack("2",2,{b},[]{"a",}) --> res{}
