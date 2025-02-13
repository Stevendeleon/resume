package pkg

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_ReadYamlFile(t *testing.T) {
	tmpDir := t.TempDir()
	dataDir := filepath.Join(tmpDir, "data")
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test data directory: %v", err)
	}

	validYaml := `
Name: John Doe
MyRole: Software Engineer
Email: john@example.com
Phone: "123-456-7890"
Location: New York, NY
PageTitle: John's Software Engineering Resume
Mission: To build amazing software
Education: BS in Computer Science

Experience:
  - Role: Senior Developer
    Company: Tech Corp
    Start: "2020"
    End: Present
    Location: New York, NY
    Details:
      - Led team of 5 developers
      - Implemented CI/CD pipeline
      - Reduced deployment time by 50%

Projects:
  - Open Source Project 1
  - Personal Website

Languages:
  - Go
  - Python
  - JavaScript

Skills:
  - Backend Development
  - Cloud Computing
  - Docker

Keywords:
  - golang
  - kubernetes
  - aws
`
	err = os.WriteFile(filepath.Join(dataDir, "valid.yml"), []byte(validYaml), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	invalidYaml := `
Name: Test
invalid yaml content
:`
	err = os.WriteFile(filepath.Join(dataDir, "invalid.yml"), []byte(invalidYaml), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid test file: %v", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(currentDir)

	tests := []struct {
		name     string
		fileName string
		want     *Resume
		wantErr  bool
	}{
		{
			name:     "valid yaml file",
			fileName: "valid",
			want: &Resume{
				Name:      "John Doe",
				Role:      "Software Engineer",
				Email:     "john@example.com",
				Phone:     "123-456-7890",
				Location:  "New York, NY",
				PageTitle: "John's Software Engineering Resume",
				Mission:   "To build amazing software",
				Education: "BS in Computer Science",
				Experience: []Experience{
					{
						Role:     "Senior Developer",
						Company:  "Tech Corp",
						Start:    "2020",
						End:      "Present",
						Location: "New York, NY",
						Details: []string{
							"Led team of 5 developers",
							"Implemented CI/CD pipeline",
							"Reduced deployment time by 50%",
						},
					},
				},
				Projects: []string{
					"Open Source Project 1",
					"Personal Website",
				},
				Languages: []string{
					"Go",
					"Python",
					"JavaScript",
				},
				Skills: []string{
					"Backend Development",
					"Cloud Computing",
					"Docker",
				},
				Keywords: []string{
					"golang",
					"kubernetes",
					"aws",
				},
			},
			wantErr: false,
		},
		{
			name:     "nonexistent file",
			fileName: "nonexistent",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid yaml content",
			fileName: "invalid",
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadYamlFile(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYamlFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadYamlFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
