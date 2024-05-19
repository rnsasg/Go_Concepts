package main

import "fmt"

func main() {
	reverseVowels("hello")
}

func reverseVowels(s string) string {

	vmap := map[byte]byte{'a': 'a', 'e': 'e', 'i': 'i', 'o': 'o', 'u': 'u', 'A': 'A', 'E': 'E', 'I': 'I', 'O': 'O', 'U': 'U'}
	i := 0
	j := len(s) - 1

	for i < j {

		for i < j {
			if _, ok := vmap[s[i]]; ok == true {
				break
			}
			i++
		}

		for i < j {
			if _, ok1 := vmap[s[j]]; ok1 == true {
				break
			}
			j--
		}
		fmt.Println(s[i], s[j])

		// if i < j {
		// 	s[i], s[j] = byte(s[j]), byte(s[i])
		// }
	}
	return s
}
