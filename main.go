package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"crawler/worker"
)

const (
	WorkerCount   = 60
	TaskQueueSize = 5000
)

// Worker
func main() {
	runtime.GOMAXPROCS(1)

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
	http.Get(u.url)
	log.Println(u.url)
}
