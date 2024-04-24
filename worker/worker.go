package worker

import "sync"

type WorkerPool interface {
	AddJob(Job)
	Start()
	Stop()
	//Output() <-chan interface{}
}

type workerPool struct {
	workerSize int
	wg         sync.WaitGroup
	workers    chan struct{}
	jobs       chan Job
}

func NewWorkers(workerSize int, queueSize int) WorkerPool {
	return &workerPool{
		workerSize: workerSize,
		wg:         sync.WaitGroup{},
		jobs:       make(chan Job, queueSize),
		workers:    make(chan struct{}, workerSize),
	}
}

func (w *workerPool) AddJob(j Job) {
	w.jobs <- j
}

func (w *workerPool) Stop() {
	close(w.jobs)
	w.wg.Wait()
}

func (w *workerPool) Start() {
	for i := 0; i < w.workerSize; i++ {
		w.wg.Add(1)
		go func() {
			for j := range w.jobs {
				w.workers <- struct{}{}
				j.Process()
				<-w.workers
			}
			w.wg.Done()
		}()
	}
}
