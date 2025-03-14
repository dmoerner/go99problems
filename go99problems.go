// Package go99problems is a set of implementations of Go solutions
// to the classic 99 Prolog problems.
package go99problems

import (
	"fmt"
	"slices"
)

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

// IsPalindrome (6): Find out whether a list is a palindrome.
func IsPalindrome[S ~[]E, E comparable](s S) bool {
	mid := len(s) / 2
	reversed := MyReverse(s[len(s)-1-mid:])
	for i := range mid {
		if s[i] != reversed[i] {
			return false
		}
	}
	return true
}

// Flatten (7): Flatten a nested list structure
//
// This has no direct parallel in the Go type system. Without ADTs, we also cannot
// construct an exact analogue with the slices and elements interspersed. The closest
// I can think of is a recursively defined struct with separate values and children.
//
// What is the efficient way to do this? slices.Concat will return a new slice
// for variadic arguments, but I don't think we can get all the other slices
// without doing a map. So we will just have to return a new slice on each recursive
// call.
type NestedSlice[T any] struct {
	values   []T
	children []*NestedSlice[T]
}

func Flatten[T any](nested NestedSlice[T]) []T {
	flattened := nested.values
	for _, child := range nested.children {
		flattened = slices.Concat(flattened, Flatten(*child))
	}
	return flattened
}

// Compress (8): Eliminate consecutive duplicates of list elements.
//
// Implemented to return a new slice.
func Compress[S ~[]E, E comparable](s S) S {
	var compressed S
	lastUnique := 0
	for i, v := range s {
		if i == 0 {
			compressed = append(compressed, v)
			continue
		}

		if v != s[lastUnique] {
			compressed = append(compressed, v)
			lastUnique = i
		}
	}
	return compressed
}

// Pack (9): Pack consecutive duplicates of list elements into sublists.
func Pack[S ~[]E, E comparable](s S) []S {
	if s == nil {
		return nil
	}

	var packed []S

	thisEl := s[0]
	var thisElSlice S
	for _, v := range s {
		if v != thisEl {
			packed = append(packed, thisElSlice)
			thisElSlice = nil

		}
		thisElSlice = append(thisElSlice, v)
		thisEl = v
	}

	packed = append(packed, thisElSlice)

	return packed
}
