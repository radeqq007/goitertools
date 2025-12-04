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
			i += step
			ch <- i
		}
	}()

	return ch
}
