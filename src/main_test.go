package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(3, 5)
	if total != 8 {
		t.Errorf("Test Failed, got: %d, want: %d.", total, 8)
	}
}
