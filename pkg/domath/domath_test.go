package domath

import "testing"

type testStruct struct {
	in1  int
	in2  int
	want int
}

func TestAddIntegers(t *testing.T) {
	cases := []testStruct{
		{1, 2, 3},
		{5, 5, 10},
		{6, 6, 12},
	}

	for _, c := range cases {
		got := AddIntegers(c.in1, c.in2)
		if got != c.want {
			t.Errorf("%q + %q != %q", c.in1, c.in2, c.want)
		}
	}
}
