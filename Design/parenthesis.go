package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var str string

	backtrack(n, 0, 0, str, &ans)
	return ans
}
func backtrack(n, open, closed int, str string, ans *[]string) {
	if open == n && closed == n {
		*ans = append(*ans, string(str))
		return
	}

	if open < n {
		backtrack(n, open+1, closed, str+"(", ans)
	}

	if closed < n && open > closed {
		backtrack(n, open, closed+1, str+")", ans)
	}
}
