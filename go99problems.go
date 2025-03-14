// Package go99problems is a set of implementations of Go solutions
// to the classic 99 Prolog problems.
package go99problems

import "fmt"

// MyLast (1): Find the last element of a list.
func MyLast[T any](s []T) (T, error) {
	if len(s) < 1 {
		return *new(T), fmt.Errorf("list must have at least one element")
	}
	return s[len(s)-1], nil
}

// MyButLast (2): Find the second-to-last element of a list.
func MyButLast[T any](s []T) (T, error) {
	if len(s) < 2 {
		return *new(T), fmt.Errorf("list must have at least two elements")
	}
	return s[len(s)-2], nil
}
