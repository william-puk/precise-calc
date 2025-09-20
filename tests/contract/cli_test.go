package contract

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIBasicUsage(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
		exitCode int
	}{
		{[]string{"5 + 3"}, "8", 0},
		{[]string{"0xFF + 1"}, "256", 0},
		{[]string{"0.1 + 0.2"}, "0.3", 0},
		{[]string{"2 + 3 x 4"}, "14", 0},
		{[]string{"-5 + 3"}, "-2", 0},
	}

	for _, test := range tests {
		// Get absolute path to binary
		workDir, _ := os.Getwd()
		binaryPath := filepath.Join(workDir, "..", "..", "bin", "precise-calc")

		cmd := exec.Command(binaryPath, test.args...)
		output, err := cmd.CombinedOutput()

		outputStr := strings.TrimSpace(string(output))

		if test.exitCode == 0 && err != nil {
			t.Errorf("Command %v expected success, got error: %v", test.args, err)
			continue
		}

		if test.exitCode != 0 && err == nil {
			t.Errorf("Command %v expected error, got success", test.args)
			continue
		}

		if !strings.Contains(outputStr, test.expected) {
			t.Errorf("Command %v output %q does not contain expected %q",
				test.args, outputStr, test.expected)
		}
	}
}

func TestCLIErrorCases(t *testing.T) {
	tests := []struct {
		args        []string
		expectedErr string
		expectFail  bool
	}{
		{[]string{"5 / 0"}, "Error", true},
		{[]string{"5 + @"}, "Error", true},
		{[]string{""}, "Error", true},
		{[]string{"5 + + 3"}, "Error", true},
		{[]string{"0xGHI + 5"}, "Error", true},
	}

	for _, test := range tests {
		// Get absolute path to binary
		workDir, _ := os.Getwd()
		binaryPath := filepath.Join(workDir, "..", "..", "bin", "precise-calc")

		cmd := exec.Command(binaryPath, test.args...)
		output, err := cmd.CombinedOutput()

		outputStr := strings.TrimSpace(string(output))

		if test.expectFail && err == nil {
			t.Errorf("Command %v expected to fail, got success", test.args)
			continue
		}

		if test.expectFail && !strings.Contains(outputStr, test.expectedErr) {
			t.Errorf("Command %v output %q does not contain expected error %q",
				test.args, outputStr, test.expectedErr)
		}
	}
}
