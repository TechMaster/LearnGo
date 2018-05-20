package work_pool

import (
	_ "expvar"
	"fmt"
	"sync"
)
type job struct {
	task		int
	wsConID		string
}


type WorkPool struct {
	maxQueueSize int
	maxWorkers int
	outputFolder string
	jobs chan job
	socketServer si.ISocketNotify
}
var wp *WorkPool //Singleton instance
var once sync.Once

func New(maxQueueSize int, maxWorkers int,  ss si.ISocketNotify) *WorkPool {
	once.Do(func() {
		wp = &WorkPool{
			maxQueueSize: maxQueueSize,
			maxWorkers: maxWorkers,
			socketServer: ss,
		}
		// create job channel
		wp.jobs = make(chan job, wp.maxQueueSize)

		// create workers
		for i := 1; i <= wp.maxWorkers; i++ {
			go func(i int) {
				for job := range wp.jobs {
					/*vd := NewVideoEncoder(wp.outputFolder, wp.socketServer)
					vd.encodeVideoToHLS(i, job)*/
				}
			}(i)
		}
	})

	return wp
}


func (es *WorkPool)EnqueueVideoEncodeRequest(filename string, wsConID string) {

	// Create Job and push the work onto the jobCh.
	job := job{filename, wsConID}
	go func() {
		fmt.Printf("added: %s %s\n", job.fileName, job.wsConID)
		es.jobs <- job
	}()
	return
}