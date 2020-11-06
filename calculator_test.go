package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {

	testCases := []testCase{
		{a: 4, b: 3, want: 7},
		{a: -7.5, b: 8, want: 0.5},
		{a: 4, b: -1, want: 3},
	}

	t.Parallel()

	for _, tc := range testCases {

		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {

	testCases := []testCase{
		{a: 4, b: 3, want: 1},
		{a: 7, b: 8, want: -1},
		{a: 4, b: 0, want: 4},
	}

	t.Parallel()
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {

	testCases := []testCase{
		{a: 4, b: 3, want: 12},
		{a: 7, b: 8, want: 56},
		{a: 4, b: 0, want: 0},
	}

	t.Parallel()

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []testCase{
		{a: 8, b: 2, want: 4, errExpected: false},
		{a: 6, b: 1.5, want: 4, errExpected: false},
		{a: 8, b: 0, want: 0, errExpected: true},
		{a: 8, b: -5, want: -1.6, errExpected: false},
	}

	t.Parallel()

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		// test for getting an error when not expecting, or no error when expecting
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived) // t.Fatalf() allows you to skip any further checks within this test (though it doesn’t stop other test functions from being run. It bails out of this test, not all tests.
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddExtreme(t *testing.T) {
	t.Parallel()
	count := 1000000
	for i := 0; i < count; i++ {
		f := rand.Float64()
		g := rand.Float64()

		r := f + g
		// fmt.Printf("Testing %f + %f = %f", f, g, r)

		tc := testCase{a: f, b: g, want: r}
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}
