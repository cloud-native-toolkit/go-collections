package go_collections

import (
	"slices"
)

// Any returns true if any element in the collection satisfies the provided predicate function, otherwise false.
func Any[S ~[]E, E any](collection S, fn func(v E) bool) bool {
	for _, value := range collection {
		if fn(value) {
			return true
		}
	}

	return false
}

// All returns true if all elements in the collection satisfy the provided predicate function, otherwise false.
func All[S ~[]E, E any](collection S, fn func(v E) bool) bool {
	if len(collection) == 0 {
		return false
	}

	for _, value := range collection {
		if !fn(value) {
			return false
		}
	}

	return true
}

// Map applies the provided function fn to each element of the input slice collection and returns a new slice with transformed elements.
func Map[S ~[]E, E any, R any](collection S, fn func(v E) R) []R {

	if len(collection) == 0 {
		return []R{}
	}

	var result []R
	for _, value := range collection {
		result = append(result, fn(value))
	}

	return result
}

// Filter iterates over the provided collection and returns a new collection containing only the elements that satisfy fn.
func Filter[S ~[]E, E any](collection S, fn func(v E) bool) S {

	deleteFn := func(v E) bool {
		return !fn(v)
	}

	return slices.DeleteFunc(collection, deleteFn)
}

// Index returns the index of the first occurrence of the test element in the collection or -1 if it is not found.
//
// Deprecated: Use slices.Index instead
func Index[S ~[]E, E comparable](collection S, test E) int {
	return slices.Index(collection, test)
}

// IndexFunc returns the index of the first element in the collection that satisfies the provided function.
// It returns -1 if no element satisfies the function.
//
// Deprecated: Use slices.IndexFunc instead
func IndexFunc[S ~[]E, E any](collection S, fn func(v E) bool) int {
	return slices.IndexFunc(collection, fn)
}

// Includes checks if a specified element exists in the given slice. Returns true if the element is found, otherwise false.
//
// Deprecated: Use slices.Contains instead
func Includes[S ~[]E, E comparable](collection S, test E) bool {
	return slices.Contains(collection, test)
}

// IncludesFunc checks if any element in the collection satisfies the specified function condition and returns true if so.
//
// Deprecated: Use slices.ContainsFunc instead
func IncludesFunc[S ~[]E, E any](collection S, fn func(v E) bool) bool {
	return slices.ContainsFunc(collection, fn)
}

// Reduce iterates over a collection, applying a reducer function to combine elements into a single result.
// The function takes a collection, a reducer function, and an initial value to accumulate the result.
func Reduce[S ~[]E, E any, R any](collection S, fn func(result R, current E, c S, index int) R, startValue R) R {
	result := startValue

	for i, value := range collection {
		result = fn(result, value, collection, i)
	}

	return result
}
