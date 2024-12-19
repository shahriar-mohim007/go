package main

//Worker Pool Pattern Breakdown
//
//A worker pool in Go allows us to create a fixed number of goroutines (workers) that process tasks (jobs) concurrently.
//	This is useful in scenarios where you have a large number of tasks to process,
//but you want to limit the number of concurrent tasks being executed, avoiding resource exhaustion and optimizing performance.

import (
	"fmt"
	"time"
)

func workerPool(numOfWorkers int, jobs <-chan int, results chan<- int) {
	for worker := 1; worker <= numOfWorkers; worker++ {
		go Worker(worker, jobs, results)
	}
}

func Worker(worker int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d picked job %d\n", worker, job)
		results <- process(job)
		fmt.Printf("Worker %d completed job %d\n", worker, job)
	}
}

func process(job int) int {
	// complete the job
	time.Sleep(time.Millisecond * 2)
	return job * 2
}

func main() {
	totalJobs := 12
	numOfWorkers := 5
	jobsQueue := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	// Start worker pool
	workerPool(numOfWorkers, jobsQueue, results)

	// Send the jobs
	for job := 0; job < totalJobs; job++ {
		jobsQueue <- job
	}

	close(jobsQueue)
	// Collect Results
	for i := 0; i < totalJobs; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println("Done.........")
}

//A worker pool in Go is a common pattern used to manage and control the number of goroutines performing concurrent tasks.
//	This approach is particularly useful when you need to limit the number of tasks running in parallel, allowing for efficient resource management.
//A race condition occurs when two or more goroutines (concurrent threads of execution) attempt to access shared data simultaneously,
//and at least one of them modifies the data. If the access and modification are not properly synchronized, the final outcome may
//be unpredictable or incorrect.
//	This can lead to bugs that are difficult to reproduce and debug because they depend on the timing of the goroutines.
