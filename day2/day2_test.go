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

func TestPresent1(t *testing.T) {
	p := Present{Length: 2, Height: 3, Width: 4}

	if p.surfaceArea() != 52 {
		t.Fatalf("Area not calculated properly: %d", p.surfaceArea())
	} else {
		t.Logf("Area: %d\n", p.surfaceArea())
	}
	if p.slackPaper() != 6 {
		t.Fatalf("Slack not calculated properly: %d", p.slackPaper())
	} else {
		t.Logf("Slack: %d\n", p.slackPaper())
	}
	if p.paperArea() != 58 {
		t.Fatalf("Total area not calculated properly: %d", p.paperArea())
	} else {
		t.Logf("Paper: %d\n", p.paperArea())
	}
	if p.ribbonLength() != 34 {
		t.Fatalf("Ribbon length not calculated properly: %d", p.ribbonLength())
	} else {
		t.Logf("Ribbon: %d\n", p.ribbonLength())
	}
}

func TestPresent2(t *testing.T) {
	p := Present{Length: 1, Height: 1, Width: 10}
	area := p.surfaceArea()
	slack := p.slackPaper()
	total := p.paperArea()
	ribbon := p.ribbonLength()

	if p.surfaceArea() != 42 {
		t.Fatalf("Area not calculated properly: %d", p.surfaceArea())
	} else {
		t.Logf("Area: %d\n", p.surfaceArea())
	}
	if p.slackPaper() != 1 {
		t.Fatalf("Slack not calculated properly: %d", p.slackPaper())
	} else {
		t.Logf("Slack: %d\n", p.slackPaper())
	}
	if p.paperArea() != 43 {
		t.Fatalf("Total area not calculated properly: %d", p.paperArea())
	} else {
		t.Logf("Paper: %d\n", p.paperArea())
	}
	if p.ribbonLength() != 14 {
		t.Fatalf("Ribbon length not calculated properly: %d", p.ribbonLength())
	} else {
		t.Logf("Ribbon: %d\n", p.ribbonLength())
	}
}
