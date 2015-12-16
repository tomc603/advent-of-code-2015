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
	"sort"
	"strconv"
	"strings"
)

type Present struct {
	Length, Width, Height int
}

func (p *Present) minSide() []int {
	sides := []int{p.Length, p.Width, p.Height}
	sort.Ints(sides)
	return sides[0:2]
}

func (p *Present) surfaceArea() int {
	// Surface area of a rectangular cuboid is the sum of the area of its sides.
	// 2*l*w + 2*w*h + 2*h*l
	return (2 * p.Length * p.Width) +
		(2 * p.Width * p.Height) +
		(2 * p.Height * p.Length)
}

func (p *Present) slackPaper() int {
	// Slack paper is the area of the smallest side of the present
	return product(p.minSide())
}

func (p *Present) ribbonLength() int {
	// Ribbon required is the sum of the sides of the smallest face plus
	// the volume of the present.
	return (2 * sum(p.minSide())) + (p.Length * p.Width * p.Height)
}

func (p *Present) paperArea() int {
	return p.surfaceArea() + p.slackPaper()
}

func sum(i []int) int {
	var total int

	for _, v := range i {
		total += v
	}
	return total
}

func product(i []int) int {
	var total int = 1

	for _, v := range i {
		total *= v
	}
	return total
}

func parsePresents(path string) (*[]Present, error) {
	var presents []Present
	var present Present // Prevent unnecessary garbage collection

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		present.Length, present.Width, present.Height = 0, 0, 0
		line := strings.Split(scanner.Text(), "x")

		// Parse each size from the split text line, convert to an int
		if i, err := strconv.Atoi(line[0]); err != nil {
			return nil, err
		} else {
			present.Length = i
		}
		if i, err := strconv.Atoi(line[1]); err != nil {
			return nil, err
		} else {
			present.Width = i
		}
		if i, err := strconv.Atoi(line[2]); err != nil {
			return nil, err
		} else {
			present.Height = i
		}
		presents = append(presents, present)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &presents, nil
}

func main() {
	var totalArea int
	var ribbonLength int

	presents, err := parsePresents("day2.data")
	if err != nil {
		fmt.Printf("Error reading present sizes.\n%s", err)
	}

	for _, p := range *presents {
		totalArea += p.paperArea()
		ribbonLength += p.ribbonLength()
	}
	fmt.Printf("Total Area: %d sq. ft.\nRibbon Length: %d\n", totalArea, ribbonLength)
}
