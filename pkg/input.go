package pkg

import (
	"flag"
	"fmt"
	"os"
)

func GetJobType() string {
	var jobType string
	flag.StringVar(&jobType, "jobType", "", "The job type to generate a resume for")
	flag.StringVar(&jobType, "j", "", "(short)The job type to generate a resume for")
	flag.Parse()

	if jobType == "" {
		jobType = "frontend"
	}

	if jobType != "frontend" && jobType != "fullstack" {
		fmt.Printf("Invalid job type: %s, Valid values include: { frontend, fullstack }\n", jobType)
		os.Exit(1)
	}

	return jobType
}
