package pkg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ReadYamlFile(fileName string) (*Resume, error) {
	filePath := fmt.Sprintf("%s.yml", fileName)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var resume Resume
	err = yaml.Unmarshal(file, &resume)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %v", err)
	}

	return &resume, nil
}
