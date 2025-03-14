// Package go99problems is a set of implementations of Go solutions
// to the classic 99 Prolog problems.
package go99problems

import "fmt"

// MyLast (1): Find the last element of a list.
func MyLast[S ~[]E, E any](s S) (E, error) {
	if len(s) < 1 {
		return *new(E), fmt.Errorf("list must have at least one element")
	}
	return s[len(s)-1], nil
}

// MyButLast (2): Find the second-to-last element of a list.
func MyButLast[S ~[]E, E any](s S) (E, error) {
	if len(s) < 2 {
		return *new(E), fmt.Errorf("list must have at least two elements")
	}
	return s[len(s)-2], nil
}

// ElementAt (3): Find the K'th element of a list, one-indexed.
func ElementAt[S ~[]E, E any](s S, k int) (E, error) {
	if k >= len(s) {
		return *new(E), fmt.Errorf("k is greater than length of list")
	}
	return s[k-1], nil
}

// MyLength (4): Find the number of elements in a list.
//
// I do not know of another way to do this which would not depend on the built-in len.
func MyLength[S ~[]E, E any](s S) int {
	return len(s)
}

// MyReverse (5): Reverse a list.
func MyReverse[S ~[]E, E any](s S) S {
	l := len(s)
	reversed := make(S, l)
	for i, v := range s {
		reversed[l-i-1] = v
	}
	return reversed
}
