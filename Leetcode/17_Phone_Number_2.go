package main

func main() {

}

var dict = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}

	helper(digits, 0, &result, "")
	return result
}

func helper(digits string, p int, result *[]string, s string) {
	if p == len(digits) {
		*result = append(*result, s)
		return
	}

	strs := dict[digits[p]]
	for i := range strs {
		helper(digits, p+1, result, s+string(strs[i]))
	}
}
