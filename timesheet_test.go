// Copyright (c) 2019. Tom Harris <mrtom.h.84@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT LICENSE
// license that can be found in the LICENSE.txt file.

package main

import (
	"testing"
)

type processChunkTestCase struct {
	start    string
	end      string
	expected int
}

func TestProcessChunk(t *testing.T) {
	tests := []processChunkTestCase{
		{"0000", "0001", 1},
		{"0000", "2359", 1439},
		{"1200", "1300", 60},
	}

	var result int

	for _, test := range tests {
		result = processChunk(test.start, test.end)
		if result != test.expected {
			t.Errorf("Expected %d mins, got %d", test.expected, result)
		}
	}
}

type parseTimeToMinsTestCase struct {
	time     string
	expected int
}

func TestParseTimeToMins(t *testing.T) {
	tests := []parseTimeToMinsTestCase{
		{"0000", 0},
		{"0120", 80},
		{"2359", 1439},
	}

	var result int

	for _, test := range tests {
		result = parseTimeToMins(test.time)
		if result != test.expected {
			t.Errorf("Expected %d mins, got %d", test.expected, result)
		}
	}
}

type parseMinsToTimeTestCase struct {
	mins     int
	expected string
}

func TestParseMinsToTime(t *testing.T) {
	tests := []parseMinsToTimeTestCase{
		{0, "0h00m"},
		{60, "1h00m"},
		{90, "1h30m"},
		{240, "4h00m"},
	}

	var result string

	for _, test := range tests {
		result = parseMinsToTime(test.mins)
		if result != test.expected {
			t.Errorf("Expected %s mins, got %s", test.expected, result)
		}
	}
}
