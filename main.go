package main

import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	_ "github.com/jackdanger/collectlinks"
)

const (
	WorkerCount   = 30
	TaskQueueSize = 5000
)

// Worker
func main() {
	runtime.GOMAXPROCS(4)

	now := time.Now()
	baseURL := "https://www.jmit.ac.in/" //"https://google.com"

	workerChan := make(chan struct{}, WorkerCount)
	taskChan := make(chan string, TaskQueueSize)
	wg := sync.WaitGroup{}
	//parsedLinks := map[string]bool{}
	//outputChan := make(chan string, TaskQueueSize)

	for i := 0; i < WorkerCount; i++ {
		go func() {
			for task := range taskChan {
				workerChan <- struct{}{}
				func() {
					defer func() {
						log.Println("done", task)
						wg.Done()
					}()

					// if parsedLinks[task] {
					// 	return
					// }
					// parsedLinks[task] = true

					//log.Println("processing", task)

					resp, err := http.Get(task)
					if err != nil {
						log.Println("[ERROR]", err)
						return
					}
					defer resp.Body.Close()

					if resp.StatusCode != 200 {
						return
					}

					// links := collectlinks.All(resp.Body)
					// for _, link := range links {
					// 	log.Println("add", link)
					// 	wg.Add(1)
					// 	taskChan <- link
					// }
				}()
				<-workerChan
			}
		}()
	}

	//log.Println("add", baseURL)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		taskChan <- baseURL
	}

	go func() {
		for range time.NewTicker(1 * time.Second).C {
			log.Println("Stats", len(workerChan), len(taskChan))
		}
	}()

	wg.Wait()

	log.Println(time.Since(now))
	// select {
	// case result, open := <-outputChan:
	// 	if !open {
	// 		return
	// 	}
	// 	log.Println(result)
	// }
}
