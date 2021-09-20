package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var phil1 phil
var phil2 phil
var phil3 phil
var phil4 phil
var phil5 phil
var phil_error phil

var phil_id int
var phil_action int
var m sync.Mutex
var amount_eating = 0

type phil struct {
	times_eaten int
	id          int
	eating      bool
	input       chan int
	output      chan int
}

func init_phil() {
	phil1 = phil{times_eaten: 0, id: 1, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil2 = phil{times_eaten: 0, id: 2, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil3 = phil{times_eaten: 0, id: 3, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil4 = phil{times_eaten: 0, id: 4, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil5 = phil{times_eaten: 0, id: 5, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil_error = phil{times_eaten: 0, id: -1, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
}

func (philo *phil) react() {
	var fork1 = get_fork_by_id(philo.id)
	var fork2 = get_fork_by_id(philo.id%5 + 1)

	for {
		m.Lock()

		//Check if adjacent forks are available
		fork1.input <- 2
		fork2.input <- 2
		fork1.get_output()
		fork2.get_output()
		isfree1 := <-fork1.output
		isfree2 := <-fork2.output

		//If both values are true -> update eating to be true and output to adjacent forks
		if isfree1 == 2390129013 && isfree2 == 2390129013 {
			philo.eating = true
			philo.times_eaten++
			fork1.being_used = true
			fork1.times_used++
			fork2.being_used = true
			fork2.times_used++
			amount_eating++
		}
		m.Unlock()
		if philo.eating {
			//wait a bit
			r := rand.Intn((500 - 300)) + 300
			time.Sleep(time.Duration(r) * time.Millisecond)

			//update eating to false
			philo.eating = false
			fork1.being_used = false
			fork2.being_used = false
			amount_eating--
		}
	}
}

func (philo *phil) phil_output() {
	action := <-philo.input
	switch action {
	case 1:
		fmt.Println("Number of times eaten: " + strconv.Itoa(philo.times_eaten))
	case 2:
		if philo.eating {
			fmt.Println("Eating")
		} else {
			fmt.Println("Thinking")
		}
	default:
		fmt.Println("Invalid input")
	}

}

func get_phil_by_id(id int) *phil {
	switch id {
	case 1:
		return &phil1
	case 2:
		return &phil2
	case 3:
		return &phil3
	case 4:
		return &phil4
	case 5:
		return &phil5
	}
	return &phil_error
}
