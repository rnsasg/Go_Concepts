package main

import "fmt"

type DivisionError struct {
	code int
	msg  string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("Code %d : %s ", d.code, d.msg)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{
			code: 2000,
			msg:  "Can not divide by zero",
		}
	}

	return a / b, nil
}

func main() {
	_, err := divide(1, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
}
