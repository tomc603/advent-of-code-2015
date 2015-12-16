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
	"strings"
)

const vowels string = "aeiou"

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

func testName(n string) bool {
	/*
	* Check the string n against the following rules. If we meet the "nice" criteria
	*  we return true. Otherwise we return false for "naughty".
	*
	*  Nice Names:
	*  * Contains >=3 vowels (aeiou). ex: aei, xazegov, aeiouaeiouaeiou
	*  * Contains immediately repeating letters. ex: xx, abcdde, aabbccdd
	*  * NOT contains forbidden strings (ab, cd, pq, xy).
	 */
	var (
		count    int
		previous rune
		repeats  bool
	)

	for _, r := range n {
		// Confirm the name doesn't contain any of the forbidden combinations of letters
		// This test is at the top to prevent paying the tax of the rest of the calculations
		// for obviously naughty names.
		if strings.Contains(string(previous)+string(r), "ab") ||
			strings.Contains(string(previous)+string(r), "cd") ||
			strings.Contains(string(previous)+string(r), "pq") ||
			strings.Contains(string(previous)+string(r), "xy") {
			return false
		}

		// We found a repeated rune, so set the repeats flag
		if previous == r {
			repeats = true
		}

		// If we find a vowel, increment count
		if strings.ContainsRune(vowels, r) {
			count++
		}
		previous = r
	}

	// If we've made it this far, there are no forbidden combinations. As long as
	// we have found repeated runes and have 3 or more vowels, this name is nice!
	if repeats && count > 2 {
		return true
	}
	return false
}

func main() {
	/*
		Process an input file of data, checking each name against the following criteria

		Nice Names:
		* Contains >=3 vowels (aeiou). ex: aei, xazegov, aeiouaeiouaeiou
		* Contains immediately repeating letters. ex: xx, abcdde, aabbccdd
		* NOT contains forbidden strings (ab, cd, pq, xy).
	*/
	var naughty, nice int

	names, err := readNames("day5.data")
	if err != nil {
		fmt.Printf("Error reading names.\n%s\n", err)
	}

	for _, name := range names {
		if testName(name) {
			nice++
		} else {
			naughty++
		}
	}
	fmt.Printf("Nice: %d\n", nice)
}
