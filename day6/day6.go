/*
   Copyright 2015 Tom Cameron

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	off uint8 = iota
	on
	toggle
)

type rect struct {
	x, y, xrange, yrange int
}

var lights [1000][1000]int

func parseInput(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if err := handleLine(scanner.Text()); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func handleLine(s string) error {
	// Split the line into a slice of commands
	cmd := strings.Split(s, " ")

	switch cmd[0] {
	case "toggle":
		r, err := handleRange(cmd[1:])
		if err != nil {
			return err
		}
		lightSet(r, toggle)
	case "turn":
		// Find the range we need to operate on early so we can simplify
		// this code block.
		r, err := handleRange(cmd[2:])
		if err != nil {
			return err
		}

		switch cmd[1] {
		case "on":
			lightSet(r, on)
		case "off":
			lightSet(r, off)
		default:
			// Only "on" or "off" should follow "turn". This indicates bad input.
			return fmt.Errorf("Unknown token %s in line %s", cmd[1], cmd)
		}
	default:
		// We should only see "turn on", "turn off", and "toggle" commands. All others
		// are unknown, and probably indicate bad input.
		return fmt.Errorf("Unknown token %s in line %s", cmd[0], cmd)
	}
	return nil
}

func handleRange(s []string) (rect, error) {
	var (
		err error
		r   rect
	)

	// Verify the string has the correct format by validating a known work exists.
	if s[1] != "through" {
		return rect{}, fmt.Errorf("Invalid range in command %s", s)
	}

	// Split the first and second X, Y coordinates from s[0] and s[2].
	// Convert the four strings into integers and set the rect{} values.
	llCorner := strings.Split(s[0], ",")
	trCorner := strings.Split(s[2], ",")

	if r.x, err = strconv.Atoi(llCorner[0]); err != nil {
		return r, err
	}
	if r.y, err = strconv.Atoi(llCorner[1]); err != nil {
		return r, err
	}

	if r.xrange, err = strconv.Atoi(trCorner[0]); err != nil {
		return r, err
	}
	if r.yrange, err = strconv.Atoi(trCorner[1]); err != nil {
		return r, err
	}

	return r, nil
}

func lightSet(r rect, v uint8) {
	for x := r.x; x <= r.xrange; x++ {
		for y := r.y; y <= r.yrange; y++ {
			switch v {
			case on:
				lights[x][y]++
			case off:
				if lights[x][y] > 0 {
					lights[x][y]--
				}
			case toggle:
				lights[x][y] += 2
			}
		}
	}
}

func main() {
	var brightness int

	if err := parseInput("day6.txt"); err != nil {
		fmt.Printf("Error parsing input file. %s\n", err)
	}
	for x := range lights {
		for y := range lights[x] {
			brightness += lights[x][y]
		}
	}
	fmt.Printf("Total brightness: %d\n", brightness)
}
