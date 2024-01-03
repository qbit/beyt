package main

import (
	"testing"
	"time"
)

func TestBeat(t *testing.T) {
	b := Beat(time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC).Unix())
	exp := 562
	if b != exp {
		t.Errorf("expected %d got %d", exp, b)
	}
}
