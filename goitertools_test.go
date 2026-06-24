package goitertools_test

import (
	"testing"

	"github.com/radeqq007/goitertools"
)

func TestCycle(t *testing.T) {
	items := []int{1, 2, 3}
	var got []int

	for v := range goitertools.Cycle(items) {
		got = append(got, v)
		if len(got) == 8 {
			break
		}
	}

	expected := []int{1, 2, 3, 1, 2, 3, 1, 2}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestCount(t *testing.T) {
	start := 10
	step := 2

	var got []int

	for v := range goitertools.Count(start, step) {
		if len(got) == 5 {
			break
		}
		got = append(got, v)
	}

	expected := []int{start, start + step, start + step*2, start + step*3, start + step*4}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("Count output mismatch at index %d: got %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestRepeat(t *testing.T) {
	val := 2
	times := 4

	got := make([]int, 0, times)
	/* 	for i := 0; i < times; i++ {
		select {
		case val := <-ch:
			got = append(got, val)
		case <-time.After(time.Second):
			t.Fatal("Repeat channel timed out")
		}
	 }*/

	for v := range goitertools.Repeat(val, times) {
		got = append(got, v)
	}

	expected := []int{2, 2, 2, 2}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestFilter(t *testing.T) {
	items := []int{10, 15, 2, 20, 5}
	condition := func(_ int, val int) bool { return val%2 == 0 }

	var got []int

	for v := range goitertools.Filter(items, condition) {
		got = append(got, v)
	}

	expected := []int{10, 2, 20}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestFilterFalse(t *testing.T) {
	items := []int{10, 15, 2, 20, 5}
	condition := func(_ int, val int) bool { return val%2 == 0 }

	var got []int

	for v := range goitertools.FilterFalse(items, condition) {
		got = append(got, v)
	}

	expected := []int{15, 5}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestCompress(t *testing.T) {
	data := []int{10, 20, 30, 40, 50}
	selectors := []bool{true, false, true, false, true}

	var got []int
	for val := range goitertools.Compress(data, selectors) {
		got = append(got, val)
	}

	expected := []int{10, 30, 50}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestDropWhile(t *testing.T) {
	items := []int{1, 3, 5, 2, 4, 6}
	condition := func(_ int, val int) bool { return val < 4 }

	var got []int
	for val := range goitertools.DropWhile(items, condition) {
		got = append(got, val)
	}

	expected := []int{5, 2, 4, 6}
	if len(got) != len(expected) {
		t.Fatalf("length mismatch: got %d elements, want %d", len(got), len(expected))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestTakeWhile(t *testing.T) {
	items := []int{2, 4, 6, 1, 4, 5}
	condition := func(_ int, val int) bool { return val%2 == 0 }

	var got []int
	for val := range goitertools.TakeWhile(items, condition) {
		got = append(got, val)
	}

	expected := []int{2, 4, 6}
	if len(got) != len(expected) {
		t.Fatalf("length mismatch: got %d elements, want %d", len(got), len(expected))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

func TestChain(t *testing.T) {
	slices := [][]int{
		{10, 2},
		{5, 1},
		{3, 1, 6},
	}

	var got []int
	for val := range goitertools.Chain(slices...) {
		got = append(got, val)
	}

	expected := []int{10, 2, 5, 1, 3, 1, 6}

	if len(got) != len(expected) {
		t.Fatalf("length mismatch: got %d elements, want %d", len(got), len(expected))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("mismatch at %d: got %d expected %d", i, got[i], expected[i])
		}
	}
}

