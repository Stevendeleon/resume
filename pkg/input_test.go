package pkg

import (
	"flag"
	"os"
	"testing"
)

func Test_GetJobType(t *testing.T) {
	originalArgs := os.Args
	defer func() {
		os.Args = originalArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "no args provided should return default frontend",
			args:     []string{"cmd"},
			expected: "frontend",
		},
		{
			name:     "short flag -j",
			args:     []string{"cmd", "-j", "backend"},
			expected: "backend",
		},
		{
			name:     "long flag --jobType",
			args:     []string{"cmd", "--jobType", "fullstack"},
			expected: "fullstack",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			os.Args = tt.args

			got := GetJobType()
			if got != tt.expected {
				t.Errorf("GetJobType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
