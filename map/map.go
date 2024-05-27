package main

import "fmt"

func main() {
	map1 := map[string]bool{"Interview": true, "Bit": true}
	map2 := map[string]bool{"Interview": true, "Questions": true}
	map3 := map1
	map1 = map2 //copy description

	fmt.Println("MAP-1", map1)
	fmt.Println("MAP-2", map2)
	fmt.Println("MAP-3", map3)
}
