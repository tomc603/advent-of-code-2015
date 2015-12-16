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
	"crypto/md5"
	"fmt"
	"strings"
)

const key string = "iwrupvqb"

// Generate md5 hashes of key+i until we find a hash whose hex value
// begins with "00000"
func main() {
	var i int

	for {
		h := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", key, i))))
		if strings.HasPrefix(h, "00000") {
			fmt.Printf("Integer %d produces a coin!\n", i)
			break
		}
		i++
	}
}
