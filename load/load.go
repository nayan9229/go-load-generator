package load

import (
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/HdrHistogram/hdrhistogram-go"
	"github.com/dustin/go-humanize"
	"github.com/jbenet/goprocess"
	"github.com/valyala/fasthttp"
)

func StartLoadTest(uri string, timeout, runtime, parallelRequests int) string {
	proc := goprocess.Background()
	clients := int(parallelRequests / 10)
	if clients < 1 {
		clients = 1
	}

	// clients := 100               // The number of connections to open to the server.
	pipeliningFactor := int(parallelRequests / clients) // The number of pipelined requests to use.
	// timeout := 20          // The number of seconds before timing out on a request
	// runtime := 60          // The number of seconds to run the autocannnon.
	debug := true // A utility debug flag.

	latencies := hdrhistogram.New(1, 10000, 5)
	requests := hdrhistogram.New(1, 1000000, 5)
	throughput := hdrhistogram.New(1, 100000000000, 5)

	var bytes int64 = 0
	var totalBytes int64 = 0
	var respCounter int64 = 0
	var totalResp int64 = 0

	resp2xx := 0
	respN2xx := 0

	errors := 0
	timeouts := 0

	ticker := time.NewTicker(time.Second)
	runTimeout := time.NewTimer(time.Second * time.Duration(runtime))

	respChan, errChan := runClients(proc, clients, pipeliningFactor, time.Second*time.Duration(timeout), uri)

	html := ``

	for {
		select {
		case err := <-errChan:
			errors++
			if debug {
				fmt.Printf("there was an error: %s\n", err.Error())
			}
			if err == fasthttp.ErrTimeout {
				timeouts++
			}
		case res := <-respChan:
			s := int64(res.size)
			bytes += s
			totalBytes += s
			respCounter++

			totalResp++
			if res.status >= 200 && res.status < 300 {
				latencies.RecordValue(int64(res.latency))
				resp2xx++
			} else {
				respN2xx++
			}

		case <-ticker.C:
			requests.RecordValue(respCounter)
			respCounter = 0
			throughput.RecordValue(bytes)
			bytes = 0
			// fmt.Println("done ticking")
		case <-runTimeout.C:
			html += getHtml()
			latencyRaw := ``
			latencyRaw += `<td>` + fmt.Sprintf("%v ms", latencies.ValueAtPercentile(2.5)) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%v ms", latencies.ValueAtPercentile(50)) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%v ms", latencies.ValueAtPercentile(97.5)) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%v ms", latencies.ValueAtPercentile(99)) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%.2f ms", latencies.Mean()) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%.2f ms", latencies.StdDev()) + `</td>`
			latencyRaw += `<td>` + fmt.Sprintf("%v ms", latencies.Max()) + `</td>`

			requestsRaw := ``
			requestsRaw += `<td>` + fmt.Sprintf("%v", requests.ValueAtPercentile(1)) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%v", requests.ValueAtPercentile(2.5)) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%v", requests.ValueAtPercentile(50)) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%v", requests.ValueAtPercentile(97.5)) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%.2f", requests.Mean()) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%.2f", requests.StdDev()) + `</td>`
			requestsRaw += `<td>` + fmt.Sprintf("%v", requests.Min()) + `</td>`

			bytesRaw := ``
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.ValueAtPercentile(1)))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.ValueAtPercentile(2.5)))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.ValueAtPercentile(50)))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.ValueAtPercentile(97.5)))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.Mean()))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.StdDev()))) + `</td>`
			bytesRaw += `<td>` + fmt.Sprintf("%v", humanize.Bytes(uint64(throughput.Min()))) + `</td>`

			html = strings.ReplaceAll(html, "[LATENCY_RAW]", latencyRaw)
			html = strings.ReplaceAll(html, "[REQ_RAW]", requestsRaw)
			html = strings.ReplaceAll(html, "[BYTES_RAW]", bytesRaw)

			summary := ``
			summary += `<p></p>`
			summary += `<p>Req/Bytes counts sampled once per second.</p>`
			summary += `<p></p>`
			summary += `<p>` + fmt.Sprintf("%v 2xx responses, %v non 2xx responses.", resp2xx, respN2xx) + `</p>`
			summary += `<p>` + fmt.Sprintf("%v total requests in %v seconds, %s read.", FormatBigNum(float64(totalResp)), runtime, humanize.Bytes(uint64(totalBytes))) + `</p>`

			if errors > 0 {
				summary += `<p>` + fmt.Sprintf("%v total errors (%v timeouts).", FormatBigNum(float64(errors)), FormatBigNum(float64(timeouts))) + `</p>`
			}
			summary += `<p>Done!</p>`
			html = strings.ReplaceAll(html, "[SUMMARY]", summary)

			return html
		}
	}
}

type resp struct {
	status  int
	latency int64
	size    int
}

func FormatBigNum(i float64) string {
	if i < 1000 {
		return fmt.Sprintf("%.0f", i)
	}
	return fmt.Sprintf("%.0fk", math.Round(i/1000))
}

func runClients(ctx goprocess.Process, clients int, pipeliningFactor int, timeout time.Duration, uri string) (<-chan *resp, <-chan error) {
	respChan := make(chan *resp, 2*clients*pipeliningFactor)
	errChan := make(chan error, 2*clients*pipeliningFactor)
	u, e := url.Parse(uri)
	if e != nil {
		fmt.Printf("e: %v\n", e)
		return respChan, errChan
	}

	for i := 0; i < clients; i++ {
		c := fasthttp.PipelineClient{
			Addr:               getAddr(u),
			IsTLS:              u.Scheme == "https",
			MaxPendingRequests: pipeliningFactor,
		}

		for j := 0; j < pipeliningFactor; j++ {
			go func() {
				req := fasthttp.AcquireRequest()
				req.SetBody([]byte("hello, world!"))
				req.SetRequestURI(uri)

				res := fasthttp.AcquireResponse()

				for {
					startTime := time.Now()
					if err := c.DoTimeout(req, res, timeout); err != nil {
						errChan <- err
					} else {
						size := len(res.Body()) + 2
						res.Header.VisitAll(func(key, value []byte) {
							size += len(key) + len(value) + 2
						})
						respChan <- &resp{
							status:  res.Header.StatusCode(),
							latency: time.Now().Sub(startTime).Milliseconds(),
							size:    size,
						}
						res.Reset()
					}
				}
			}()
		}
	}
	return respChan, errChan
}

// getAddr returns the address from a URL, including the port if it's not empty.
// So it can return hostname:port or simply hostname
func getAddr(u *url.URL) string {
	if u.Port() == "" {
		return u.Hostname()
	} else {
		return fmt.Sprintf("%v:%v", u.Hostname(), u.Port())
	}
}

func getHtml() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Report</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
			}
			table {
				width: 100%;
				border-collapse: collapse;
				margin-bottom: 20px;
			}
			th, td {
				padding: 12px;
				text-align: left;
				border: 1px solid #ddd;
			}
			th {
				background-color: #f4f4f4;
			}
			tbody tr:nth-child(even) {
				background-color: #f9f9f9;
			}
			p {
				margin: 10px 0;
			}
		</style>
	</head>
	<body>
		<table>
			<thead>
				<tr>
					<th>Stat</th>
					<th>2.5</th>
					<th>50</th>
					<th>97.5</th>
					<th>99</th>
					<th>Avg</th>
					<th>Stdev</th>
					<th>Max</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Latency</td>
					[LATENCY_RAW]
				</tr>
			</tbody>
		</table>
		<table>
			<thead>
				<tr>
					<th>Stat</th>
					<th>1</th>
					<th>2.5</th>
					<th>50</th>
					<th>97.5</th>
					<th>Avg</th>
					<th>Stdev</th>
					<th>Min</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Req/Sec</td>
					[REQ_RAW]
				</tr>
				<tr>
					<td>Bytes/Sec</td>
					[BYTES_RAW]
				</tr>
			</tbody>
		</table>
		[SUMMARY]
	</body>
	</html>
	`
}
