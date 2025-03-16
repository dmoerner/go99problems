// Package go99problems is a set of implementations of Go solutions
// to the classic 99 Prolog problems.
package go99problems

import (
	"fmt"
	"math/rand"
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

// Encode (10): Run-length encoding of a list.
//
// Since Go does not have tuples, we need to build a generic struct. Without a built-in
// or efficient `map` function, it makes sense to rewrite this rather than using Pack.
type RLEPair[E comparable] struct {
	Count   int
	Element E
}

func Encode[S ~[]E, E comparable](s S) []RLEPair[E] {
	if s == nil {
		return nil
	}

	var encoded []RLEPair[E]

	thisEl := s[0]
	thisCount := 0
	for _, v := range s {
		if v != thisEl {
			encoded = append(encoded, RLEPair[E]{thisCount, thisEl})
			thisCount = 0

		}
		thisCount += 1
		thisEl = v
	}

	encoded = append(encoded, RLEPair[E]{thisCount, thisEl})

	return encoded
}

// EncodeModified (11): This does not make much sense in Go because slices are
// all of one type unless you use an interface slice, which should be avoided
// at all costs. Other languages like Haskell with the same contraints use option
// ADTs for this, which Go also does not support. The struct approach from Problem
// 10 is best, but here's a simple interface implementation anyway with a flat slice.
func EncodeModified[S ~[]E, E comparable](s S) [][]any {
	if s == nil {
		return nil
	}

	var encoded [][]any

	thisEl := s[0]
	thisCount := 0
	for _, v := range s {
		if v != thisEl {
			encoded = append(encoded, []any{thisCount, thisEl})
			thisCount = 0
		}
		thisCount += 1
		thisEl = v
	}

	encoded = append(encoded, []any{thisCount, thisEl})

	return encoded
}

// Decode (12): Decode a run-length encoded list. The original versions use the
// modified encode from Problem 11, but we will use the more idiomatic struct
// implementation I made for Problem 10.
func Decode[E comparable](e []RLEPair[E]) []E {
	var decoded []E

	for _, v := range e {
		for range v.Count {
			decoded = append(decoded, v.Element)
		}
	}

	return decoded
}

// EncodeDirect (13): RLE, direct solution
//
// I find the spec for this somewhat unclear, but I believe the task is
// to edit the encoding in-place, rather than appending a completed tuple
// to the list.
func EncodeDirect[S ~[]E, E comparable](s S) []RLEPair[E] {
	if s == nil {
		return nil
	}

	encoded := []RLEPair[E]{{1, s[0]}}

	for _, v := range s[1:] {
		if v == encoded[len(encoded)-1].Element {
			encoded[len(encoded)-1].Count += 1
		} else {
			encoded = append(encoded, RLEPair[E]{1, v})
		}
	}

	return encoded
}

// Dupli (14): Duplicate the elements of a list.
func Dupli[S ~[]E, E any](s S) S {
	var duplicated S
	for _, v := range s {
		duplicated = append(duplicated, v)
		duplicated = append(duplicated, v)
	}
	return duplicated
}

// Repli (15): Replicate the elements of a list a given number of times.
func Repli[S ~[]E, E any](s S, count int) S {
	var replicated S
	for _, v := range s {
		for range count {
			replicated = append(replicated, v)
		}
	}
	return replicated
}

// DropEvery (16): Drop every N'th element from a list.
func DropEvery[S ~[]E, E any](s S, n int) S {
	var dropped S
	for i, v := range s {
		if (i+1)%n == 0 {
			continue
		}
		dropped = append(dropped, v)
	}
	return dropped
}

// Split (17): Split a list into two parts, the length of the first part is
// given. To continue our theme, we return copies.
func Split[S ~[]E, E any](s S, n int) (S, S) {
	left := make(S, n)
	right := make(S, len(s)-n)

	copy(left, s[:n])
	copy(right, s[n:])

	return left, right
}

// Slice (18): Extract a slice from a list. Inclusive range, 1-indexed.
// Return a copy.
func Slice[S ~[]E, E any](s S, i int, k int) S {
	sliced := make(S, k-i+1)
	copy(sliced, s[i-1:k+1])
	return sliced
}

// Rotate (19): Rotate a list N places to the left.
func Rotate[S ~[]E, E any](s S, n int) S {
	places := n % len(s)
	if places < 0 {
		places += len(s)
	}
	return slices.Concat(s[places:], s[:places])
}

// RemoveAt (20): Remove the K'th element from a list. 1-indexed.
// Return a new list.
func RemoveAt[S ~[]E, E any](k int, s S) (E, S) {
	if k < 1 || k > len(s) {
		return *new(E), s
	}

	return s[k-1], slices.Concat(s[:k-1], s[k:])
}

// InsertAt (21): Insert an element at a given position into a list.
// 1-indexed, return a new list.
func InsertAt[S ~[]E, E any](el E, s S, k int) S {
	if k < 1 {
		k = 1
	} else if k > len(s) {
		k = len(s)
	}

	newslice := make(S, 0, len(s)+1)

	newslice = append(newslice, s[:k-1]...)
	newslice = append(newslice, el)
	newslice = append(newslice, s[k-1:]...)

	return newslice
}

// Range (22): Create a list containing all integers within a given range, inclusive.
func Range(start, end int) []int {
	newslice := make([]int, 0, end-start+1)
	for i := range end - start + 1 {
		newslice = append(newslice, start+i)
	}
	return newslice
}

// RndSelect (23): Extract a given number of randomly selected elements
// from a list. The specification does not say if duplicates are prohibited,
// so we allow them.
func RndSelect[S ~[]E, E any](s S, count int) S {
	if count < 1 {
		return *new(S)
	}

	selected := make(S, 0, count)

	for range count {
		selected = append(selected, s[rand.Intn(len(s))])
	}

	return selected
}

// RndSelectUnique: Variation on Problem 23 which does not pick the same
// element twice. How do we do this efficiently? My technique is to shuffle
// all possible indices and then take a subset of those. I wish this were
// lazier, but it should be far more efficient than manually checking for
// dupes.
func RndSelectUnique[S ~[]E, E any](s S, count int) S {
	if count < 1 {
		return *new(S)
	}

	if count > len(s) {
		count = len(s)
	}

	selected := make(S, 0, count)

	indices := Range(0, len(s)-1)

	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	for i := range count {
		selected = append(selected, s[indices[i]])
	}

	return selected
}

// DiffSelect (24): Draw N different random numbers from the set 1..M.
func DiffSelect(n, m int) []int {
	possibilities := Range(1, m)

	selected := RndSelectUnique(possibilities, n)

	return selected
}
