package main

import (
	"testing"
)

func TestProcessData(t *testing.T) {
	var tests = []struct {
		description string
		port        int
		expected    int
	}{
		{"PortWithinRange", 8080, 0},
		{"PortOutsideRange", 80800, 1},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			errCode := realMain(test.port, true)
			if test.expected != errCode {
				t.Errorf("Unexpected result %d", errCode)
			}
		})
	}
}
