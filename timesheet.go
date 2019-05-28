// Copyright 2019 Tom Harris <mrtom.h.84@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT LICENSE
// license that can be found in the LICENSE.txt file.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	times := os.Args[1:]
	if len(times)%2 > 0 {
		log.Fatal("Incorrect number of arguments")
	}

	mins := 0

	for i := 0; i < len(times); i += 2 {
		mins += processChunk(times[i], times[i+1])
	}

	fmt.Print(parseMinsToTime(mins) + "\n")
}

func processChunk(start string, end string) int {
	return parseTimeToMins(end) - parseTimeToMins(start)
}

// Turn mins into a formatted string representing
// hours and mins. e.g. mins: 60 would result in the string
// "1h00m" being returned.
func parseMinsToTime(mins int) string {
	hours := mins / 60
	minutes := mins % 60

	// Format the result as, for example, 1h30m or 0h10m or 12h00m
	return fmt.Sprintf("%dh%02dm", hours, minutes)
}

func parseTimeToMins(time string) int {
	h, m := parseSubTime(time[:2]), parseSubTime(time[2:])

	return m + (h * 60)
}

func parseSubTime(time string) int {
	intTime, err := strconv.Atoi(time)
	if err != nil {
		log.Fatalf("Not integer %s", time)
	}

	return intTime
}
