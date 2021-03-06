// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package datasrcs

import (
	"testing"
)

func TestBaidu(t *testing.T) {
	if *networkTest == false {
		return
	}

	result := testDNSRequest("Baidu")
	if result < expectedTest {
		t.Errorf("Found %d names, expected at least %d instead", result, expectedTest)
	}
}
