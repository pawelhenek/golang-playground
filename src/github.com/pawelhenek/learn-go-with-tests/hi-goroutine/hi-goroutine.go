package main

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // send 'i' to channel 'ch'.
	}
}

func Filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // loop over values received
		if i%prime != 0 {
			dst <- i // send 'i' to channel 'dst'.
		}
	}
}

func main() {
	src := make(chan int) // create a new channel
	go Generate(src)
	for i := 0; i < 100; i++ { // find 100 primes
		prime := <-src
		println(prime)
		dst := make(chan int)
		go Filter(src, dst, prime)
		src = dst
	}
}
