package main

import (
	"runtime"
	"sync"
)

type mutedSlice struct {
	mux sync.Mutex
	data []int
	done []bool
}

func mutedSliceConstructor(size int) *mutedSlice {
    this := new(mutedSlice)
    this.data = make([]int, size)
    this.done = make([]bool, size)
    return this
}

func reader(f func(int) int, in <-chan int, data *mutedSlice, n int)  {

	for counter := 0; counter < n; counter++ {
		x := <- in

		go func(place int, val int) {
			defer data.mux.Unlock()

			res := f(val)
			data.mux.Lock()
			data.data[place] = res
			data.done[place] = true
			return
		} (counter, x)
	}
}

func checker(out chan<- int, data1 *mutedSlice, data2 *mutedSlice, n int) {

	go func() {
		for counter := 0; counter < n; {
			data1.mux.Lock()
			data2.mux.Lock()

			for counter < n && data1.done[counter] && data2.done[counter] {
				out <- data1.data[counter] + data2.data[counter]
				counter ++
			}

			data1.mux.Unlock()
			data2.mux.Unlock()

			runtime.Gosched()
		}
	} ()
}

func sender(in <-chan int, out chan<- int, n int) {
	for counter := 0; counter < n; counter++ {
		out <- <-in
	}
}

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int)  {

	go func() {
		pipe := make(chan int, n)

		data1 := mutedSliceConstructor(n)
		data2 := mutedSliceConstructor(n)

		go reader(f, in1, data1, n)
		go reader(f, in2, data2, n)

		go checker(pipe, data1, data2, n)

		go sender(pipe, out, n)
	} ()
}
