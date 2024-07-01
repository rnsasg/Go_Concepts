package main

import "fmt"

type myString string

type alias = string

func main() {
	var str1 myString = "This is my string of type myString"
	var str2 alias = "It is alias of string type"

	var copy1 string = str2
	var copy2 string = str1 // cannot use str1 (variable of type myString) as string value in variable declaration

	fmt.Println(copy1, copy2)

}

// func main() {
// 	var str1 myString = "This is my string of type myString"
// 	var str2 alias = "It is alias of string type"
// 	fmt.Printf("%T - %s\n", str1, str1)
// 	fmt.Printf("%T - %s\n", str2, str2)
// }

// Output :
// go run other/defined_vs_alias.go
// main.myString - This is my string of type myString
// string - It is alias of string type
