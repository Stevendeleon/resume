package pkg

import (
	"flag"
)

// provide the file as a flag -j or -jobType <file-name>
// this file should live in /data as a yaml
func GetJobType() string {
	var jobType string

	flag.StringVar(&jobType, "jobType", "", "The job type to generate a resume for")
	flag.StringVar(&jobType, "j", "", "(short)The job type to generate a resume for")
	flag.Parse()

	if jobType == "" {
		jobType = "frontend"
	}

	return jobType
}
