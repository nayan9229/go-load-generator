package handlers

import (
	"net/http"
	"strconv"

	"github.com/nayan9229/go-load-generator/db"
	"github.com/nayan9229/go-load-generator/model"
	"github.com/nayan9229/go-load-generator/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index(db.GetJobsInstance().Jobs.All()))
}

func HandleHomePost(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	url := r.Form.Get("url")
	if url == "" {
		return Render(w, r, home.Index(db.GetJobsInstance().Jobs.All()))
	}
	urlTimeoutStr := r.Form.Get("url_timeout")
	urlTimeout, _ := strconv.Atoi(urlTimeoutStr)
	runtimeStr := r.Form.Get("runtime")
	parallelRequestsStr := r.Form.Get("parallel_requests")
	runtime, _ := strconv.Atoi(runtimeStr)
	parallelRequests, _ := strconv.Atoi(parallelRequestsStr)

	db.GetJobsInstance().Jobs.Add(url, urlTimeout, runtime, parallelRequests, model.Pending)
	return Render(w, r, home.Index(db.GetJobsInstance().Jobs.All()))
}
