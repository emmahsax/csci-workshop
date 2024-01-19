package main

import "testing"

// Format method starting with "Test" and then any character that isn't a lowercase letter

// func Test_isPrime(t *testing.T) {
// 	n := 0
// 	result, msg := isPrime(n)
// 	if result {
// 		t.Errorf("with %d as test parameter, got true, but expected false", n)
// 	}

// 	if msg != "0 is not prime, by definition!" {
// 		t.Error("wrong message returned:", msg)
// 	}

// 	n = 7
// 	result, msg = isPrime(n)
// 	if !result {
// 		t.Errorf("with %d as test parameter, got false, but expected true", n)
// 	}

// 	if msg != "7 is a prime number!" {
// 		t.Error("wrong message returned:", msg)
// 	}
// }

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"negative number", -3, false, "-3 is negative, and negative numbers are not prime, by definition!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true, but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false, but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s, but  got %s", e.name, e.msg, msg)
		}
	}
}
