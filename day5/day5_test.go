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

import "testing"

func TestNames(t *testing.T) {
	// Test "nice" cases
	if !checkName("qjhvhtzxzqqjkmpb") {
		t.Fatal("Nice name returned naughty.\n")
	} else {
		t.Log("Nice name OK.")
	}

	if !checkName("qjhvhtzxzqqqjkmpb") {
		t.Fatal("Nice name returned naughty.\n")
	} else {
		t.Log("Nice name OK.")
	}

	if !checkName("xxyxx") {
		t.Fatal("Nice name returned naughty.\n")
	} else {
		t.Log("Nice name OK.")
	}

	// Test "naughty" cases
	if checkName("aaa") {
		t.Fatal("Naughty name returned nice.\n")
	} else {
		t.Log("Naughty name OK.")
	}

	if checkName("uurcxstgmygtbstg") {
		t.Fatal("Naughty name returned nice.\n")
	} else {
		t.Log("Naughty name OK.")
	}

	if checkName("ieodomkazucvgmuy") {
		t.Fatal("Naughty name returned nice.\n")
	} else {
		t.Log("Naughty name OK.")
	}
}
