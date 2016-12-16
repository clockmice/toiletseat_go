package main

import "testing"

func TestPolicy1(t *testing.T) {
	m := make(map[string]int)
	m["UUUDDUDU"] = 6
	m["DDUUDD"] = 5
	m["UDUDUDU"] = 6

	for k, v := range m {
		expected := v
		actual := policy1(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}

func TestPolicy2(t *testing.T) {
	m := make(map[string]int)
	m["UUUDDUDU"] = 7
	m["DDUUDD"] = 4
	m["UDUDUDU"] = 7

	for k, v := range m {
		expected := v
		actual := policy2(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}

func TestPolicy3(t *testing.T) {
	m := make(map[string]int)
	m["UUUDDUDU"] = 4
	m["DDUUDD"] = 2
	m["UDUDUDU"] = 6

	for k, v := range m {
		expected := v
		actual := policy3(k)
		if actual != expected {
			t.Error("Test1 failed")
		}
	}
}