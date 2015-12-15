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

func parseInput(path string) ([]string, error) {
	var results []string

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		break
	}
	return results, nil
}

func main() {
	for i, c := range parseInput("dayX.data") {
		fmt.Printf("%d: %s", i, c)
	}
}
