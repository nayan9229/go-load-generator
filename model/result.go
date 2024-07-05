package model

import "github.com/HdrHistogram/hdrhistogram-go"

type Result struct {
	Latencies   *hdrhistogram.Histogram
	Requests    *hdrhistogram.Histogram
	Throughput  *hdrhistogram.Histogram
	Bytes       int64
	TotalBytes  int64
	RespCounter int64
	TotalResp   int64
	Resp2xx     int
	RespN2xx    int
	Errors      int
	Timeouts    int
}

func NewResult() *Result {
	return &Result{
		Latencies:   hdrhistogram.New(1, 10000, 5),
		Requests:    hdrhistogram.New(1, 1000000, 5),
		Throughput:  hdrhistogram.New(1, 100000000000, 5),
		Bytes:       0,
		TotalBytes:  0,
		RespCounter: 0,
		TotalResp:   0,
		Resp2xx:     0,
		RespN2xx:    0,
		Errors:      0,
		Timeouts:    0,
	}
}
