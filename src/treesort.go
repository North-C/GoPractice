package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// append the elements of t to values in order
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	v := [...]int{28, 343, 38, 10, 3490, 38, 290}

	for _, value := range v {
		fmt.Printf("%d ", value)
	}
	fmt.Println(" ")

	s := Sort(v[:])
	for _, value := range s {
		fmt.Printf("%d ", value)
	}
	fmt.Println(" ")

}
