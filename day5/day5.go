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
)

func readNames(path string) ([]string, error) {
	var names []string

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s.\n%s\n", path, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	for s.Scan() {
		names = append(names, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return names, nil
}

func cheapRing(s []rune, r rune) []rune {
	// Simulate a rune ring buffer so we always have current and two prior for comparison
	if len(s) > 2 {
		s = append(s[1:], r)
	} else {
		s = append(s, r)
	}
	return s
}

func checkName(n string) bool {
	/*
	* Check the string n against the following rules. If we meet the "nice" criteria
	*  we return true. Otherwise we return false for "naughty".
	*
	*  Nice Names:
	*  * Contains a pair of repeated runes that do not overlap. ex: xyxy, aabcdefgaa, not aaa
	*  * Contains at least one letter that is repeated with one letter between them. ex: xyx, aaa
	 */
	var (
		triplet bool
		ring    []rune
	)

	// fmt.Printf("\nName: %s\n", n)
	pairs := map[string]int{}
	for _, r := range n {
		ring = cheapRing(ring, r)

		// Minimum sensible ring size is 2 elements
		if len(ring) == 2 {
			pairs[string(ring[0])+string(ring[1])]++
		} else if len(ring) > 2 {
			if ring[0] == ring[2] {
				// A nice triplet contains a repeated letter at the beginning and end
				triplet = true
				// fmt.Printf("Adding triplet: %s\n", string(ring))
			}

			// The current and previous rune pair needs to be counted because it is not
			// an overlapping pair]
			if ring[1] != ring[2] {
				// fmt.Printf("Adding pair: %s from %s\n", string(ring[1:]), string(ring))
				pairs[string(ring[1:])]++
			} else if ring[0] != ring[2] || ring[0] != ring[1] {
				// fmt.Printf("Adding pair: %s from %s\n", string(ring[1:]), string(ring))
				pairs[string(ring[1:])]++
			}
		}
	}
	// fmt.Printf("Pairs: %v\n", pairs)

	// fmt.Printf("Name: %s\nPairs: %v\n", n, pairs)
	// If we have a repeated letter, and a pair repeated more than once the name is nice.
	if triplet {
		// No sense in iterating through pairs if we don't have a pepeated letter
		for _, v := range pairs {
			if v > 1 {
				return true
			}
		}
	}
	return false
}

func main() {
	// Process an input file of data, checking each name against the following criteria
	var naughty, nice int

	names, err := readNames("day5.data")
	if err != nil {
		fmt.Printf("Error reading names.\n%s\n", err)
	}

	for _, name := range names {
		if checkName(name) {
			nice++
		} else {
			naughty++
		}
	}
	fmt.Printf("Nice: %d\n", nice)
}
