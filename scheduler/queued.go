package scheduler

import (
	"crawler/engine"
	"log"
	"time"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	wokerChan   chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.wokerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.wokerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		t := time.Tick(5 * time.Second)
		var requestQ []engine.Request
		var WorkerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(WorkerQ) > 0 {
				activeWorker = WorkerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.wokerChan:
				WorkerQ = append(WorkerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				WorkerQ = WorkerQ[1:]
			case <-t:
				log.Printf("requestQ: %v, WorkerQ: %v\n", len(requestQ), len(WorkerQ))
			}
		}
	}()
}
