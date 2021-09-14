package main

import "fmt"

var input = make(chan int, 1)  //receives information from philosophers next to this
var output = make(chan int, 1) //send to philosophers next to this
var fork1 fork
var fork2 fork
var fork3 fork
var fork4 fork
var fork5 fork

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
	//input <- 2
	//fmt.Println(fork1.get_output(input))
}

func (f fork) react(nr int) {
	if nr == true_nr {
		f.being_used = true
	}
	fmt.Println(f.being_used)
}

func (f fork) get_output(actions <-chan int) int {

	num := <-actions
	//fmt.Println(num)
	if num == 2 {
		if f.being_used {
			return false_nr
		} else {
			return true_nr
		}
	}

	return 0

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
	fmt.Errorf("Error: id out of scope")
	return &fork1
}
