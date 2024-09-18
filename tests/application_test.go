package tests

import (
	"os"
	"os/exec"
	"testing"
)

const (
	executable = "./geektrust"
	soureDir   = ".."
)

func TestGeektrustEndToEnd(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want string
	}{}

	// Build binary
	cmd := exec.Command("go", "build", "-o", executable, soureDir)
	if buildOutput, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %s", buildOutput)
	}

	defer os.Remove(executable)

	// Execute tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command(executable, test.args.input)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf(
					"application execution failed error= %s, output= %q",
					err,
					string(output),
				)
			}

			if got := string(output); got != test.want {
				t.Errorf("./geektrust = %q, want= %q", got, test.want)
			}
		})
	}
}
