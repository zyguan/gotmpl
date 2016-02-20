package gomap

import "sync"

// template type GoMap(A, B)

type A int

type B int

func GoMap(n int, f func(A) B, in <-chan A) <-chan B {
	if n <= 0 {
		n = 1
	}
	out := make(chan B)
	go func() {
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				for a := range in {
					out <- f(a)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(out)
	}()
	return out
}
