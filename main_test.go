package main

import "testing"

func TestPolicy1(t *testing.T) {
	m := map[string]int {
		"UUUDDUDU": 6,
		"DDUUDD": 5,
		"UDUDUDU": 6,
	}

	for k, v := range m {
		expected := v
		actual := policy1(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}

func TestPolicy2(t *testing.T) {
	m := map[string]int {
		"UUUDDUDU": 7,
		"DDUUDD": 4,
		"UDUDUDU": 7,
	}

	for k, v := range m {
		expected := v
		actual := policy2(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}

func TestPolicy3(t *testing.T) {
	m := map[string]int {
		"UUUDDUDU": 4,
		"DDUUDD": 2,
		"UDUDUDU": 6,
	}

	for k, v := range m {
		expected := v
		actual := policy3(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}