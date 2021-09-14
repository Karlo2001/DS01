package main

import "fmt"

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

	//Take input while the goroutines run
	for {
		fmt.Println("\nType 1 to query forks")
		fmt.Println("Type 2 to query philosophers")
		fmt.Scanln(&fork_or_philo)
		if fork_or_philo == 2 {
			fmt.Println("Insert the id of the philosopher to query (1-5)")
			fmt.Scanln(&phil_id)
			selected_phil = *get_phil_by_id(phil_id)
			if selected_phil.id != -1 {
				fmt.Println("Type 1 for nr times eaten")
				fmt.Println("Type 2 for eating or thinking")
				fmt.Scanln(&phil_action)
				selected_phil.phil_output(phil_action)
			} else {
				fmt.Println("Selected ID is not valid")
			}
		} else {
			fmt.Println("Insert the id of the fork to query (1-5)")
			fmt.Scanln(&fork_id)
			selected_fork = *get_fork_by_id(fork_id)
			if selected_fork.id != -1 {
				fmt.Println("Type 1 for nr times used")
				fmt.Println("Type 2 for being used")
				fmt.Scanln(&fork_action)
				selected_fork.fork_output(fork_action)
			}
		}
	}
}
