package counter

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		ctr    counter
		expect int
	}{
		{
			ctr:    counter{c: []int{1, 2, 3, 0}},
			expect: 6,
		},
		{
			ctr:    counter{c: []int{1, 2, 3, 50}},
			expect: 56,
		},
	}
	for _, c := range cases {
		result := c.ctr.Count()
		if c.expect != result {
			t.Errorf("Expected sum to be \"%d\" but get \"%d\" ", c.expect, result)
		}
	}
}

func TestCountSecond(t *testing.T) {
	ctr := &counter{
		c: []int{1, 2, 3, 0},
	}

	count := ctr.Count()

	expect := 6
	if count != expect {
		t.Errorf("fatal: expect sum to be \"%d\" but get \"%d\"\n", expect, count)
	}
}

func TestCountAdd(t *testing.T) {
	initArr := []int{1, 2, 3, 0}
	expectArr := []int{2, 3, 0, 0}

	ctr := &counter{
		c: initArr[:],
	}

	ctr.Add()

	if !reflect.DeepEqual(ctr.c, expectArr) {
		t.Errorf("fatal: expect arr to be \"%+v\" but get \"%+v\"\n", expectArr, ctr.c)
	}
}
