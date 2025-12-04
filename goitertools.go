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
