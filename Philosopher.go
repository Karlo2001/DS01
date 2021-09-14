package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//var input = make(chan int, 1) //receives information from forks next to this
//var output = make(chan int, 1) //return information to be read elsewhere
var m sync.Mutex

type phil struct {
	times_eaten   int
	using_fork_id int
	id            int
	eating        bool
	input         chan int
	output        chan int
}

func init_phil() {
	phil1 := phil{times_eaten: 0, id: 1, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil2 := phil{times_eaten: 0, id: 2, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil3 := phil{times_eaten: 0, id: 3, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil4 := phil{times_eaten: 0, id: 4, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	phil5 := phil{times_eaten: 0, id: 5, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	phil1.react()
	phil5.react()
}

func (philo phil) react() {
	m.Lock()
	//Check if adjacent forks are available
	var fork1 = get_fork_by_id(philo.id)
	var fork2 = get_fork_by_id(philo.id%5 + 1)

	fork1.input <- 2
	fork2.input <- 2
	isfree1 := fork1.get_output(fork1.input)
	isfree2 := fork2.get_output(fork2.input)

	fmt.Println(isfree1)
	//If both values are true -> update eating to be true and output to adjacent forks
	if isfree1 == true_nr && isfree2 == true_nr {
		philo.eating = true
		philo.times_eaten++
		fork1.being_used = true
		fork1.times_used++
		fork2.being_used = true
		fork2.times_used++
	}

	if philo.eating {
		//fork1.input <- I am eating
		//fork2.input <- I am eating
		//wait a bit
		r := rand.Intn(10)
		time.Sleep(time.Duration(r) * time.Microsecond)

		//update eating to false
		philo.eating = false

		//fork1.input <- I am not eating
		//fork2.input <- I am not eating
	}

	//Test
	fmt.Println(fork1)
	fmt.Println(fork2)
	m.Unlock()
}
