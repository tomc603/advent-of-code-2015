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
	"io"
	"os"
	"strings"
)

type House struct {
	x, y int
}

func parseDirections(path string) (map[House]int, error) {
	var (
		house House // Prevent unnecessary garbage collection
		x, y  int
	)
	m := make(map[House]int)
	m[house]++ // Initial delivery to origin point

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		} else {
			switch {
			case strings.ContainsRune("^", c):
				x++
			case strings.ContainsRune("v", c):
				x--
			case strings.ContainsRune("<", c):
				y--
			case strings.ContainsRune(">", c):
				y++
			}
			house.x, house.y = x, y
		}
		m[house]++
	}
	return m, nil
}

func main() {
	// Parse the input file and follow the directions to each house.
	// The x, y coordinates of each house are a key in a map. The value is the
	// integer count of times we've visited.
	var deliveries int
	houses, err := parseDirections("day3.data")

	if err != nil {
		fmt.Printf("ERROR: Could not follow the directions!\n%s\n", err)
	} else {
		for _, v := range houses {
			if v >= 1 {
				deliveries++
			}
		}
		fmt.Printf("Deliveries: %d\n", deliveries)
	}
}
