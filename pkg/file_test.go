package pkg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadYamlFile_FileNotFound(t *testing.T) {
	_, err := ReadYamlFile("non_existing_file")
	if err == nil {
		t.Errorf("Expected error for non-existing file, got nil")
	}
}

func TestReadYamlFile_InvalidYaml(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_invalid_yaml_file")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err = os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	invalidYamlContent := `: invalid yaml`
	fileName := filepath.Join(tempDir, "invalid_resume.yml")
	err = os.WriteFile(fileName, []byte(invalidYamlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ReadYamlFile(fileName[:len(fileName)-4])
	if err == nil {
		t.Errorf("Expected error for invalid YAML, got nil")
	}
}
