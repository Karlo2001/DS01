package main

import (
	"sync"
)

var fork1 fork
var fork2 fork
var fork3 fork
var fork4 fork
var fork5 fork
var fork_error fork

type fork struct {
	times_used int
	id         int
	being_used bool
	input      chan int
	output     chan int
	fm         sync.Mutex
}

func init_forks() {
	fork1 = fork{times_used: 0, id: 1, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork2 = fork{times_used: 0, id: 2, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork3 = fork{times_used: 0, id: 3, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork4 = fork{times_used: 0, id: 4, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork5 = fork{times_used: 0, id: 5, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork_error = fork{times_used: 0, id: -1, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
}

func (f *fork) act() {
	for {
		f.check_input()
	}
}

func (f *fork) check_input() {
	select {
	case in := <-f.input:
		if in == 1 {
			f.output <- f.times_used
		} else if in == 2 {
			if f.being_used {
				f.output <- true_nr
			} else {
				f.output <- false_nr
			}
		}
	default:
	}
}

func get_fork_by_id(nr int) *fork {
	if nr == 1 {
		return &fork1
	}
	if nr == 2 {
		return &fork2
	}
	if nr == 3 {
		return &fork3
	}
	if nr == 4 {
		return &fork4
	}
	if nr == 5 {
		return &fork5
	}
	return &fork_error
}
