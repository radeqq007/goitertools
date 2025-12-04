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
