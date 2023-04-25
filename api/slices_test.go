package api

import (
	"fmt"
	"testing"
)

func TestReverse0(t *testing.T) {
	x := []string{"1", "2", "3"}
	fmt.Printf("%v\n", x)
	count := len(x)
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	fmt.Printf("%v\n", x)
}

func TestReverse(t *testing.T) {
	x := []string{"1", "2", "3", "4"}
	fmt.Printf("%v\n", x)
	xk := Reverse(x)
	fmt.Printf("%v\n", xk)
}

func TestReverse1(t *testing.T) {
	x := [5]string{"1", "2", "3", "4", "6"}
	fmt.Printf("%v\n", x)
	xk := Reverse(x[:])
	fmt.Printf("%v\n", xk)
}

func Remove1(slice []interface{}, start, end int) []interface{} {
	result := make([]interface{}, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

// --------------------另一种更为简便的写法-----------------------
func Remove2(slice []interface{}, start, end int) []interface{} {
	return append(slice[:start], slice[end:]...)
}

func TestFilter(t *testing.T) {
	x := []string{"9", "1", "2", "3", "4", "6"}
	fmt.Printf("%v\n", x)
	xk := Filter(x[:], func(s string) bool {
		return s > "3"
	})
	fmt.Printf("%v\n", xk)
}
