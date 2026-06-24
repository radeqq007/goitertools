// Package goitertools provides a collection iterator utilities 
// inspired by Python's itertools library
package goitertools

import "iter"

// Cycle returns an infinite iterator that repeatedly yields elements from the slice.
// It stops immediately if items is empty.
func Cycle[T any](items []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		if len(items) == 0 {
			return
		}

		for {
			for _, item := range items {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Count returns an infinite iterator that yields evenly spaced values starting.
// from the 'start', incrementing by 'step' on each iteration.
func Count(start, step int) iter.Seq[int] {
	return func(yield func(int) bool) {
		cur := start
		for {
			if !yield(cur) {
				return
			}

			cur += step
		}
	}
}

// Repeat returns an iterator that yields the specified value a fixed number of times.
// If 'times' is less than 1, the iterator yields nothing.
func Repeat[T any](value T, times int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range times {
			if !yield(value) {
				return
			}
		}
	}
}

// Filter returns an iterator that yields elements from the slice for which the condition function returns true.
// The condition function receives both the index and the value of the element.
func Filter[T any](items []T, condition func(int, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, v := range items {
			if condition(i, v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// FilterFalse returns an iterator that yields elements from the slice for which the condition function returns false.
// The condition function receives both the index and the value of the element.
func FilterFalse[T any](items []T, condition func(int, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, v := range items {
			if !condition(i, v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Compress returns an iterator that filters elements returning only those that have a corresponding true value in the selectors.
// Iteration stops when either the 'items' or 'selectors' slice is exhausted.
func Compress[T any](items []T, selectors []bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		minLen := len(items)
		minLen = min(minLen, len(selectors))

		for i := range minLen {
			if selectors[i] {
				if !yield(items[i]) {
					return
				}
			}
		}
	}
}

// DropWhile returns an iterator that skips elements as long as the condition function returns true.
// Once the condition returns false, the and all remaining elements are yielded without further evaluation of the condition.
func DropWhile[T any](items []T, condition func(int, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		dropping := true
		for i, val := range items {
			if dropping && condition(i, val) {
				continue
			}
			dropping = false
			if !yield(val) {
				return
			}
		}
	}
}

// TakeWhile returns an iterator that yields elements from the slice as long as the condition function returns true.
// It terminates immediately once the condition returns false.
func TakeWhile[T any](items []T, condition func(int, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, val := range items {
			if condition(i, val) {
				if !yield(val) {
					return
				}
			} else {
				break
			}
		}
	}
}

// Chain returns an iterator that treats consecutive slices as a single continuous sequence.
// It yields all elements from the first slice, then the second, and so on, until all provided slices are exhausted.
func Chain[T any](slices ...[]T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, slice := range slices {
			for _, val := range slice {
				if !yield(val) {
					return
				}
			}
		}
	}
}
