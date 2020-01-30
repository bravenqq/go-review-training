package main

import (
	"testing"
	"time"
)

func TestIsRecentMonth(t *testing.T) {
	tests := []struct {
		input  time.Time
		output bool
	}{

		{input: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), output: true},
		{input: time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local), output: false},
	}

	for _, test := range tests {
		res := IsRecentMonth(test.input)
		if res != test.output {
			t.Errorf("inout:%v output:%t want:%t\n", test.input, res, test.output)
		}
	}
}

func TestIsRecentYear(t *testing.T) {
	tests := []struct {
		input  time.Time
		output bool
	}{

		{input: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), output: true},
		{input: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local), output: false},
	}

	for _, test := range tests {
		res := IsRecentYear(test.input)
		if res != test.output {
			t.Errorf("inout:%v output:%t want:%t\n", test.input, res, test.output)
		}
	}
}
