package main

import (
	"fmt"
	"os"
)

var true_nr int = 2390129013
var false_nr int = 2329132910
var selected_phil phil
var fork_id int
var fork_or_philo int
var selected_fork fork
var fork_action int

func main() {
	init_forks()
	init_phil()

	go phil1.react()
	go phil2.react()
	go phil3.react()
	go phil4.react()
	go phil5.react()
	go fork1.act()
	go fork2.act()
	go fork3.act()
	go fork4.act()
	go fork5.act()

	//Take input while the goroutines run
	for {
		fmt.Println("\nType 1 to query forks")
		fmt.Println("Type 2 to query philosophers")
		fmt.Println("Type 0 to exit program")
		fmt.Scanln(&fork_or_philo)
		if fork_or_philo == 2 {
			fmt.Println("Insert the id of the philosopher to query (1-5)")
			fmt.Scanln(&phil_id)
			selected_phil = *get_phil_by_id(phil_id)
			if selected_phil.id != -1 {
				fmt.Println("Type 1 for nr times eaten")
				fmt.Println("Type 2 for eating or thinking")
				fmt.Scanln(&phil_action)

				phil_query()
				//selected_phil.phil_output(phil_action)
			} else {
				fmt.Println("Selected ID is not valid")
			}
		} else if fork_or_philo == 1 {
			fmt.Println("Insert the id of the fork to query (1-5)")
			fmt.Scanln(&fork_id)
			selected_fork = *get_fork_by_id(fork_id)
			if selected_fork.id != -1 {
				fmt.Println("Type 1 for nr times used")
				fmt.Println("Type 2 for being used")
				fmt.Scanln(&fork_action)

				fork_query()
				//selected_fork.fork_output(fork_action)
			}
		} else if fork_or_philo == 0 {
			os.Exit(0)
		}
	}
}

func phil_query() {
	selected_phil.input <- phil_action

	if phil_action == 1 {
		fmt.Println(<-selected_phil.output)
	} else if phil_action == 2 {
		out := <-selected_phil.output
		if out == true_nr {
			fmt.Println("Eating")
		} else if out == false_nr {
			fmt.Println("Thinking")
		}
	} else {
		fmt.Println("Invalid input")
	}
}

func fork_query() {
	selected_fork.input <- fork_action

	if fork_action == 1 {
		fmt.Println(<-selected_fork.output)
	} else if fork_action == 2 {
		out := <-selected_fork.output
		if out == true_nr {
			fmt.Println("In use")
		} else if out == false_nr {
			fmt.Println("Not in use")
		}
	} else {
		fmt.Println("Invalid input")
	}
}
