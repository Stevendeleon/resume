package main

import (
	"fmt"
	"os"

	"github.com/Stevendeleon/pkg"
)

func main() {
	jobType := pkg.GetJobType()

	resume, err := pkg.ReadYamlFile(jobType)
	if err != nil {
		fmt.Println("Error parsing YAML file: ", err)
		os.Exit(1)
	}

	err = pkg.GenerateHTML(resume)
	if err != nil {
		fmt.Println("Error generating HTML: ", err)
		os.Exit(1)
	}
}
