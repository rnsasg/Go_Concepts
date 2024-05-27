package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	strs := []string{"/home/"} //
	//strs := []string{"/home/user/Documents/../Pictures"}
	// strs := []string{"/.../a/../b/c/../d/./"}
	//strs := []string{"/../"}
	//strs := []string{"/home//foo/"}
	for _, s := range strs {
		fmt.Println(simplifyPath(s))
	}
}

func simplifyPath(path string) string {
	var strs []string
	var stack *list.List = list.New()

	exp := regexp.MustCompile(`\/+`)

	fmt.Println(len(exp.Split(path, -1)))
	for _, s := range exp.Split(path, -1) {
		if s == ".." {
			if stack.Len() != 0 {
				stack.Remove(stack.Front())
			}
			continue
		} else if s == "." || s == "" {
			continue
		}
		stack.PushFront(s)
	}
	for stack.Len() != 0 {
		strs = append(strs, stack.Front().Value.(string))
		stack.Remove(stack.Front())
	}
	fmt.Println(strs)
	i := 0
	j := len(strs) - 1
	for i < j {
		strs[i], strs[j] = strs[j], strs[i]
		i++
		j--
	}

	fmt.Println(strs)
	// res := strings.Join(strs, "/")
	// return res[:len(res)-1]
	return "/" + strings.Join(strs, "/")
}
