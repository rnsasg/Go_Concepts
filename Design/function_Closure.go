package main

func main() {

}

func combine(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(int)
	dfs = func(i int) {
		if len(tmp) == k {
			copied := make([]int, k)
			copy(copied, tmp)
			res = append(res, copied)
			return
		}
		if n-i+1 < k-len(tmp) {
			return
		}
		for j := i; j <= n; j++ {
			tmp = append(tmp, j)
			dfs(j + 1)
			tmp = tmp[0 : len(tmp)-1]
		}
	}
	dfs(1)
	return res
}
