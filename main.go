package main

import (
	"log"
	"time"

	"worker-model/worker"
)

const (
	WorkerCount   = 120
	TaskQueueSize = 5000
)

// Worker
func main() {

	now := time.Now()
	baseURL := "https://www.jmit.ac.in/"

	worker := worker.NewWorkers(WorkerCount, TaskQueueSize)

	worker.Start()

	for i := 0; i < 1000; i++ {
		worker.AddJob(URLJOb{
			url: baseURL,
		})
	}

	worker.Stop()

	log.Println(time.Since(now))
}

type URLJOb struct {
	url string
}

func (u URLJOb) Process() {
	time.Sleep(5 * time.Second)
}
