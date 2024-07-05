package db

import (
	"sync"

	"github.com/nayan9229/go-load-generator/model"
)

func init() {
	// Initialize the database connection
	// for i := 0; i < 10; i++ {
	// 	jobs := GetJobsInstance()
	// 	s := model.Pending
	// 	if i%2 == 0 {
	// 		s = model.Running
	// 	}
	// 	if i%3 == 0 {
	// 		s = model.Completed
	// 	}
	// 	if i%4 == 0 {
	// 		s = model.Failed
	// 	}
	// 	jobs.Jobs.Add(fmt.Sprintf("https://%d.com", i), 10, 30, 20, s)
	// }
}

type JobDb struct {
	Jobs *model.Jobs
}

var instance *JobDb
var once sync.Once

// GetInstance returns the single instance of the Singleton struct.
func GetJobsInstance() *JobDb {
	once.Do(func() {
		jd := model.NewJobs()
		instance = &JobDb{Jobs: jd} // <-- Here we are initializing the Jobs struct
	})
	return instance
}
