package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nayan9229/go-load-generator/db"
	"github.com/nayan9229/go-load-generator/views/job_details"
)

func HandleJobDetails(w http.ResponseWriter, r *http.Request) error {
	jobID := chi.URLParam(r, "job_id")
	_ = jobID
	fmt.Printf("jobID: %v\n", jobID)
	return Render(w, r, job_details.JobDetails(db.GetJobsInstance().Jobs.Get(uuid.MustParse(jobID))))
}
