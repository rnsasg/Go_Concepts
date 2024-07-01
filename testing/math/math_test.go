package math_test

import (
	"testing"

	math "github.com/rnsasg/Go_Concepts/testing/math"
)

type addTestcase struct {
	a, b, expected int
}

var testcases = []addTestcase{
	{1, 1, 2},
	{3, 4, 7},
	{25, 25, 51},
}

func TestAdd(t *testing.T) {
	for _, tc := range testcases {
		got := math.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Expected %d but got %d", tc.expected, got)
			t.Fail()
		}
	}
}

// package math_test

// import (
// 	"testing"

// 	math "github.com/rnsasg/Go_Concepts/testing/math"
// )

// func TestAdd(t *testing.T) {
// 	res := math.Add(1, 2)
// 	expected := 3

// 	if res != expected {
// 		t.Fail()
// 	}
// }
