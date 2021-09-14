package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//var input = make(chan int, 1) //receives information from forks next to this
//var output = make(chan int, 1) //return information to be read elsewhere
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
	phil1 := phil{times_eaten: 0, id: 1, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil2 := phil{times_eaten: 0, id: 2, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil3 := phil{times_eaten: 0, id: 3, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil4 := phil{times_eaten: 0, id: 4, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	phil5 := phil{times_eaten: 0, id: 5, eating: false, input: make(chan int, 1), output: make(chan int, 1)}
	//Adding go infront of the functions makes it such that they don't go into the function?

	go phil1.react()
	go phil2.react()
	go phil3.react()
	go phil4.react()
	go phil5.react()

	//Wait while the goroutines run
	for {

	}
	//fmt.Println("Done")
}

func (philo phil) react() {
	var fork1 = get_fork_by_id(philo.id)
	var fork2 = get_fork_by_id(philo.id%5 + 1)

	for {
		m.Lock()

		//Check if adjacent forks are available
		fork1.input <- 2
		fork2.input <- 2
		isfree1 := fork1.get_output(fork1.input)
		isfree2 := fork2.get_output(fork2.input)

		//If both values are true -> update eating to be true and output to adjacent forks
		if isfree1 == 2390129013 && isfree2 == 2390129013 {
			philo.eating = true
			philo.times_eaten++
			fork1.being_used = true
			fork1.times_used++
			fork2.being_used = true
			fork2.times_used++
			amount_eating++
			fmt.Println("Philo id: " + strconv.Itoa(philo.id) + "\t times eaten: " + strconv.Itoa(philo.times_eaten) + "\t amount eating: " + strconv.Itoa(amount_eating))
		}
		m.Unlock()
		if philo.eating {
			//fork1.input <- I am eating
			//fork2.input <- I am eating
			//wait a bit
			r := rand.Intn((500 - 300)) + 300
			time.Sleep(time.Duration(r) * time.Millisecond)

			//update eating to false
			philo.eating = false
			fork1.being_used = false
			fork2.being_used = false
			amount_eating--

			//fork1.input <- I am not eating
			//fork2.input <- I am not eating
		}

		//Test
		//fmt.Println("fork1 ", fork1)
		//fmt.Println("fork2 ", fork2)
	}
}
