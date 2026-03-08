package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
