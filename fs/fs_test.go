package fs

import (
	"reflect"
	"testing"
)

func TestSliceToString(t *testing.T) {
	s := service{
		delimeter: " ",
	}
	cases := []struct {
		slc    []int
		expect string
	}{
		{
			slc:    []int{1, 2, 3},
			expect: "1 2 3",
		},
		{
			slc:    []int{1, 2, 3, 50},
			expect: "1 2 3 50",
		},
	}

	for _, c := range cases {
		result := s.sliceToString(c.slc)
		if c.expect != result {
			t.Fatalf("Expected string to ve \"%s\" but get \"%s\" ", c.expect, result)
		}
	}
}

func TestStringToSlice(t *testing.T) {
	s := service{
		delimeter: " ",
	}
	cases := []struct {
		slc    string
		expect []int
	}{
		{
			slc:    "1 2 3",
			expect: []int{1, 2, 3},
		},
		{
			slc:    "1 2 3 50",
			expect: []int{1, 2, 3, 50},
		},
	}

	for _, c := range cases {
		result := s.stringToSlice(c.slc)
		if !reflect.DeepEqual(c.expect, result) {
			t.Fatalf("Expected string to ve \"%+v\" but get \"%+v\" ", c.expect, result)
		}
	}
}
