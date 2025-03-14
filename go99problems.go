// Package go99problems is a set of implementations of Go solutions
// to the classic 99 Prolog problems.
package go99problems

// MyLast (1): Find the last element of a list.
func MyLast[T any](s []T) T {
	return s[len(s)-1]
}
