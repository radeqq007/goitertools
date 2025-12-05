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

	expected := []int{0, 0 + step, 0 + step*2, 0 + step*3, 0 + step*4}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("Count output mismatch at index %d: got %d, want %d", i, got[i], expected[i])
		}
	}
}

func TestRepeat(t *testing.T) {
	val := 0
	times := 4
	ch := goitertools.Repeat(val, times)

	got := make([]int, 0, times)
	for i := 0; i < times; i++ {
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

func TestFilter(t *testing.T) {
	items := []int{10, 15, 2, 20, 5}
	condition := func(val, _ int) bool { return val%2 == 0 }
	ch := goitertools.Filter(items, condition)

	got := []int{}
	for val := range ch {
		got = append(got, val)
	}

	expected := []int{10, 2, 20}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}

func TestFilterFalse(t *testing.T) {
	items := []int{11, 15, 2, 20, 5}
	condition := func(val, _ int) bool { return val%2 == 0 }
	ch := goitertools.FilterFalse(items, condition)

	got := []int{}
	for val := range ch {
		got = append(got, val)
	}

	expected := []int{11, 15, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}

func TestCompress(t *testing.T) {
	data := []int{10, 20, 30, 40, 50}
	selectors := []bool{true, false, true, false, true}
	ch := goitertools.Compress(data, selectors)
	
	got := []int{}
	for val := range ch {
		got = append(got, val)
	}
	
	expected := []int{10, 30, 50}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}

func TestChain(t *testing.T) {
	slices := [][]int{
		{10, 2},
		{5, 1},
		{3, 1, 6},
	}

	ch := goitertools.Chain(slices)

	got := []int{}
	for val := range ch {
		got = append(got, val...)
	}

	expected := []int{10, 2, 5, 1, 3, 1, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}

func TestChainFromSlice(t *testing.T) {
	slices := [][]int{
		{10, 2},
		{5, 1},
		{3, 1, 6},
	}

	ch := goitertools.ChainFromSlice(slices)

	got := []int{}
	for val := range ch {
		got = append(got, val)
	}

	expected := []int{10, 2, 5, 1, 3, 1, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}
