package load

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jbenet/goprocess"
	"github.com/nayan9229/go-load-generator/db"
	"github.com/nayan9229/go-load-generator/model"
	"github.com/valyala/fasthttp"
)

func StartJobProcessor() {
	for {
		job := db.GetJobsInstance().Jobs.Next()
		if job == nil {
			slog.Info("No jobs to process. Sleeping for 30 seconds")
			time.Sleep(30 * time.Second)
			continue
		}
		processJob(job)
	}
}

func processJob(job *model.Job) {
	slog.Info("Processing job", "job", job)
	doLoadTest(job)
	newJob := db.GetJobsInstance().Jobs.Next()
	if newJob != nil {
		processJob(newJob)
	}
}

func doLoadTest(job *model.Job) error {
	proc := goprocess.Background()
	result := model.NewResult()
	clients := int(job.ParallelRequests / int(float64(job.ParallelRequests)*0.2))
	if clients < 1 {
		clients = 1
	}
	pipeliningFactor := int(job.ParallelRequests / clients) // The number of pipelined requests to use.

	ticker := time.NewTicker(time.Second)
	runTimeout := time.NewTimer(time.Second * time.Duration(job.Runtime))

	respChan, errChan := runClients(proc, clients, pipeliningFactor, time.Second*time.Duration(job.Timeout), job.Uri)

	for {
		select {
		case err := <-errChan:
			result.Errors++
			if job.Debug {
				fmt.Printf("there was an error: %s\n", err.Error())
			}
			if err == fasthttp.ErrTimeout {
				result.Timeouts++
			}
		case res := <-respChan:
			s := int64(res.size)
			result.Bytes += s
			result.TotalBytes += s
			result.RespCounter++

			result.TotalResp++
			if res.status >= 200 && res.status < 300 {
				result.Latencies.RecordValue(int64(res.latency))
				result.Resp2xx++
			} else {
				result.RespN2xx++
			}

		case <-ticker.C:
			result.Requests.RecordValue(result.RespCounter)
			result.RespCounter = 0
			result.Throughput.RecordValue(result.Bytes)
			result.Bytes = 0
			// fmt.Println("done ticking")
		case <-runTimeout.C:
			job.Status = model.Completed
			job.Result = result
			return nil
		}
	}
}
