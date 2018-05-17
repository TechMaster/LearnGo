package encode_server

import (
	_ "expvar"
	"fmt"
	"sync"
	si "github.com/TechMaster/LearnGo/UploadConvertFeedback/socket_interface"
)
type job struct {
	fileName	string
	wsConID		string
}


type EncodeServer struct {
	maxQueueSize int
	maxWorkers int
	outputFolder string
	jobs chan job
	socketServer si.ISocketNotify
}
var es *EncodeServer  //Singleton instance
var once sync.Once

func GetEncodeServer() *EncodeServer {
	return es
}
func New(maxQueueSize int, maxWorkers int, outputFolder string, ss si.ISocketNotify) *EncodeServer{
	once.Do(func() {
		es = &EncodeServer{
			maxQueueSize: maxQueueSize,
			maxWorkers: maxWorkers,
			outputFolder: outputFolder,
			socketServer: ss,
		}
		// create job channel
		es.jobs = make(chan job, es.maxQueueSize)

		// create workers
		for i := 1; i <= es.maxWorkers; i++ {
			go func(i int) {
				for job := range es.jobs {
					vd := NewVideoEncoder(es.outputFolder, es.socketServer)
					//vd.encodeVideoToHLS(i, job)  //Real encoding
					vd.mockEncodeVideoHLS(i, job)
				}
			}(i)
		}
	})

	return es
}


func (es *EncodeServer)EnqueueVideoEncodeRequest(filename string, wsConID string) {

	// Create Job and push the work onto the jobCh.
	job := job{filename, wsConID}
	go func() {
		fmt.Printf("added: %s %s\n", job.fileName, job.wsConID)
		es.jobs <- job
	}()
	return
}