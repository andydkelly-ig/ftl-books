package calculator_test

import (
	"calculator"
	"testing"
	"math/rand"
	"time"
)

func TestAddS(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.AddOriginal(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestAddVariadic(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want float64
	}
	testCases := []testCase{
		{inputs: []float64{1,2,3,4}, want: 10},
		{inputs: []float64{0,1,0,0}, want: 1},
		{inputs: []float64{-10,-2,-3}, want: -15},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < 50; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 1, b: 3, want: 2},
		{a: 9, b: 7, want: -2},
		{a: -2, b: -1, want: 1},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		name string
	}
	testCases := []testCase{
		{a: 6, b: 3, want: 18, name: "Two positive integers"},
		{a: -2, b: -2, want: 4, name: "Two negatives make a positive"},
		{a: 0, b: -1, want: 0, name: "Zero multiplication = 0"},
		{a: -2, b: 2, want: -4, name: "Negative and Positive = Negative"},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f for test %s", tc.want, got, tc.name)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		name        string
		errExpected bool
	}
	testCases := []testCase{
		{a: 6, b: 3, want: 2, name: "Basic Division", errExpected: false},
		{a: 6, b: 0, want: 0, name: "Divison by zero", errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("want %f, got %f for test %s", tc.want, got, tc.name)
		}
	}
}

func TestSqrRoot(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a	 		float64
		want		float64
		name		string
		errExpected	bool
	}
	testCases := []testCase {
		{a: 4, want: 2, name: "Basic Square Root", errExpected: false},
		{a: -1, want: 999, name: "Negative Sqaure Root", errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.SqrRoot(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Sqaure Root(%f): unexpected error status %v", tc.a, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Want %f, got %f for test %s", tc.want, got, tc.name)
		}
	}
}