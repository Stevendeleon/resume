package pkg

import (
	"flag"
	"os"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestGetJobType_Default(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	resetFlags()

	os.Args = []string{"cmd"}

	jobType := GetJobType()
	expected := "frontend"
	if jobType != expected {
		t.Errorf("GetJobType() = %s, want %s", jobType, expected)
	}
}

func TestGetJobType_Fullstack(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	resetFlags()

	os.Args = []string{"cmd", "-jobType", "fullstack"}

	jobType := GetJobType()
	expected := "fullstack"
	if jobType != expected {
		t.Errorf("GetJobType() = %s, want %s", jobType, expected)
	}
}

func TestGetJobType_ShortFlag(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	resetFlags()

	os.Args = []string{"cmd", "-j", "frontend"}

	jobType := GetJobType()
	expected := "frontend"
	if jobType != expected {
		t.Errorf("GetJobType() = %s, want %s", jobType, expected)
	}
}
