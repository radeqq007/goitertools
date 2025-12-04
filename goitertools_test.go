package goitertools_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/radeqq007/goitertools"
)

func TestCycle(t *testing.T) {
	items := []int{1, 2, 3}
	ch := goitertools.Cycle(items)

	got := make([]int, 0, 6)
	for range 6 {
		select {
		case val := <-ch:
			got = append(got, val)
		case <-time.After(time.Second):
			t.Fatal("Cycle channel timed out")
		}
	}

	expected := []int{1, 2, 3, 1, 2, 3}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("Cycle output mismatch at index %d: got %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestCount(t *testing.T) {
	start := 0
	step := 2
	ch := goitertools.Count(start, step)

	got := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			got = append(got, val)
		case <-time.After(time.Second):
			t.Fatal("Count channel timed out")
		}
	}

	expected := []int{0 + step, 0 + step*2, 0 + step*3, 0 + step*4, 0 + step*5}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("Count output mismatch at index %d: got %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestRepeat(t *testing.T) {
	val := 0
	times := 4
	ch := goitertools.Count(val, times)

	got := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		select {
		case val := <-ch:
			got = append(got, val)
		case <-time.After(time.Second):
			t.Fatal("Repeat channel timed out")
		}
	}

	expected := []int{0, 0, 0, 0}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}
