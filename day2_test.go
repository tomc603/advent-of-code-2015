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

func TestArea1(t *testing.T) {
	p := Present{Length: 2, Height: 3, Width: 4}
	area := p.area()
	slack := p.slack()
	total := p.area() + p.slack()

	if area != 52 {
		t.Fatalf("Area not calculated properly: %d", area)
	}
	if slack != 6 {
		t.Fatalf("Slack not calculated properly: %d", slack)
	}
	if total != 58 {
		t.Fatalf("Total area not calculated properly: %d", total)
	}
	t.Logf("Success: %d, %d, %d", area, slack, total)
}

func TestArea2(t *testing.T) {
	p := Present{Length: 1, Height: 1, Width: 10}
	area := p.area()
	slack := p.slack()
	total := p.area() + p.slack()

	if area != 42 {
		t.Fatalf("Area not calculated properly: %d", area)
	}
	if slack != 1 {
		t.Fatalf("Slack not calculated properly: %d", slack)
	}
	if total != 43 {
		t.Fatalf("Total area not calculated properly: %d", total)
	}
	t.Logf("Success: %d, %d, %d", area, slack, total)

}
