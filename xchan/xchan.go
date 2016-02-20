package xchan

// template type XChan(A)

type A int

func XChan(xs []A) <-chan A {
	out := make(chan A)
	go func() {
		for _, x := range xs {
			out <- x
		}
		close(out)
	}()
	return out
}
