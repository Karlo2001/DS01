package main

import "fmt"

//var input = make(chan int, 1) //receives information from forks next to this
//var output = make(chan int, 1) //send to forks next to this

type phil struct {
	times_eaten int
	using_fork_id int
	id          int
	eating      bool
	input       chan int
	output      chan int
}

func init_phil() {
	phil1 := phil{times_eaten: 0, id: 1, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil2 := phil{times_eaten: 0, id: 2, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil3 := phil{times_eaten: 0, id: 3, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil4 := phil{times_eaten: 0, id: 4, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	//phil5 := phil{times_eaten: 0, id: 5, eating: false, input: make(chan int, 1), output: make(chan int, 1), using_fork_id: 0}
	phil1.react()
}

func (philo phil) react() {

	//Check if ascending forks are available
	fork1 := get_fork_by_id(philo.id)
	var fork2
	if philo.id == 5 {
		fork2 = get_fork_by_id(1)
	}
	else fork2 = get_fork_by_id(philo.id+1)
	
	philo.input := <- 
	philo.input := <- 
	//Ask forks whether they are free

	isfree1 := <- fork1.output
	isfree2 := <- fork2.output
	if isfree1 == 
	//If both values are true -> update eating to be true and output to ascending forks
	
	if philo.eating {
		//fork1.input <- I am eating 
		//fork2.input <- I am eating
		//wait a bit
		//update eating to false
		//fork1.input <- I am not eating
		//fork2.input <- I am not eating
	}

	//Test
	fmt.Println(philo.id)
}
