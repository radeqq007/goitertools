package goitertools

func Cycle[T any](items []T) <-chan T {
	ch := make(chan T)

	go func() {
		for {
			for _, item := range items {
				ch <- item
			}
		}
	}()

	return ch
}

func Count(start, step int) <-chan int {
	ch := make(chan int)

	go func() {
		i := start
		for {
			ch <- i
			i += step
		}
	}()

	return ch
}

func Repeat[T any](value T, times int) <-chan T {
	ch := make(chan T)
	go func() {
		if times == -1 {
			for {
				ch <- value
			}
		} else {
			for i := 0; i < times; i++ {
				ch <- value
			}
			close(ch)
		}
	}()

	return ch
}

func Filter[T any](items []T, condition func(T, int) bool) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)

		for i, val := range items {
			if condition(val, i) {
				ch <- val
			}
		}
	}()
	return ch
}

func FilterFalse[T any](items []T, condition func(T, int) bool) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)

		for i, val := range items {
			if !condition(val, i) {
				ch <- val
			}
		}
	}()
	return ch
}

func Compress[T any](items []T, selectors []bool) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		
		minLen := len(items)
		if len(selectors) < minLen {
			minLen = len(selectors)
		}
		
		for i := 0; i < minLen; i++ {
			if selectors[i] {
				ch <- items[i]
			}
		}
	}()
	return ch
}

func DropWhile[T any](items []T, condition func(T, int) bool) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)

		dropping := true
		for i, val := range items {
			if dropping && condition(val, i) {
				continue
			} else {
				dropping = false
				ch <- val
			}
		}
	}()

	return ch
}

func Chain[T any](slices ...[]T) <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		for _, slice := range slices {
			for _, val := range slice {
				ch <- val
			}
		}
	}()

	return ch
}

func ChainFromSlice[T any](slices [][]T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, slice := range slices {
			for _, val := range slice {
				ch <- val
			}
		}
	}()

	return ch
}
