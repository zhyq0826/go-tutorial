package main

import "time"

func (workerPool *WorkerPool) processWaitState() {
	select {
	case finished := <-workerPool.finishedChannel:
		payload := &payload{
			workerID: finished,
		}
		workerPool.goToState(stateFinish, payload)

	case processing := <-workerPool.processingChannel:
		payload := &payload{
			isProcessing: processing,
		}
		workerPool.goToState(stateProcessing, payload)
	case <-workerPool.shutdown.initiateChannel:
		workerPool.goToState(stateQuit, nil)
	case <-time.After(time.Second * time.Duration(workerPool.timeout)):
		workerPool.goToState(stateTimeout, nil)
	}
}
