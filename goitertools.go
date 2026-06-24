package goitertools

import "iter"

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

func Repeat[T any](value T, times int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range times {
			if !yield(value) {
				return
			}
		}
	}
}

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
