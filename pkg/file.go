package pkg

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed data/*.yml
var yamlFiles embed.FS

func ReadYamlFile(fileName string) (*Resume, error) {
	filePath := fmt.Sprintf("data/%s.yml", fileName)

	file, err := yamlFiles.ReadFile(filePath)
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
