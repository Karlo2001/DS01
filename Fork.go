package main

import (
	"fmt"
	"strconv"
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
}

func init_forks() {
	fork1 = fork{times_used: 0, id: 1, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork2 = fork{times_used: 0, id: 2, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork3 = fork{times_used: 0, id: 3, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork4 = fork{times_used: 0, id: 4, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork5 = fork{times_used: 0, id: 5, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
	fork_error = fork{times_used: 0, id: -1, being_used: false, input: make(chan int, 1), output: make(chan int, 1)}
}

func (f *fork) react(nr int) {
	if nr == true_nr {
		f.being_used = true
	}
	fmt.Println(f.being_used)
}

func (f *fork) get_output() {

	num := <-f.input
	if num == 1 {
		f.output <- f.times_used
	}
	if num == 2 {
		if f.being_used {
			f.output <- false_nr
		} else {
			f.output <- true_nr
		}
	}
}

func (f *fork) fork_output() {
	action := <-f.input
	switch action {
	case 1:
		fmt.Println("Number of times used: " + strconv.Itoa(f.times_used))
	case 2:
		if f.being_used {
			fmt.Println("Being used")
		} else {
			fmt.Println("Not being used")
		}
	default:
		fmt.Println("Invalid input")
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
