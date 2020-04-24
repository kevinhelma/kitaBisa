package main

import "testing"

func Test_Add(t *testing.T){
	total := add(5,5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func Test_Multiply(t *testing.T){
	total := multiply(5,5)
	if total != 25 {
		t.Errorf("Multiplication was incorrect, got: %d, want: %d.", total, 25)
	}
}

func Test_Fibonnaci(t *testing.T){
	type fibTest struct {
		n        int
		expected int
	}
	var fibTests = []fibTest {
		{1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {7, 13},
	}

	for _, tt := range fibTests {
		actual := fibonacci(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}

}

