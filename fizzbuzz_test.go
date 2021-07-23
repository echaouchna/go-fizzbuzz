package main

import (
	"testing"
)

func TestFizzbuzzValues(t *testing.T) {
	value := getFizzbuzzValue(100, FizzbuzzOptions{2, 5, "ggg", "hhh"})
	if value != "ggghhh" {
		t.Errorf("Expected: ggghhh\ngot: %q", value)
	}
	value = getFizzbuzzValue(100, FizzbuzzOptions{3, 3, "ggg", "hhh"})
	if value != "100" {
		t.Errorf("Expected: 100\ngot: %q", value)
	}
}

func TestDefaultFizzbuzzResult(t *testing.T) {
	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz,"
	actual := fizzbuzz(100, DefaultFizzbuzzOptions)
	if actual != expected {
		t.Errorf("Expected: %q\ngot: %q", expected, actual)
	}
}

func TestCustomFizzbuzzResult(t *testing.T) {
	options := FizzbuzzOptions{
		Int1: 3,
		Int2: 3,
		Str1: "fiz",
		Str2: "buz",
	}
	expected := "1,2,fizbuz,4,5,fizbuz,7,8,fizbuz,10,11,fizbuz,13,14,fizbuz,16,17,fizbuz,19,20,fizbuz,22,23,fizbuz,25,26,fizbuz,28,29,fizbuz,31,32,fizbuz,34,35,fizbuz,37,38,fizbuz,40,41,fizbuz,43,44,fizbuz,46,47,fizbuz,49,50,fizbuz,52,53,fizbuz,55,56,fizbuz,58,59,fizbuz,61,62,fizbuz,64,65,fizbuz,67,68,fizbuz,70,71,fizbuz,73,74,fizbuz,76,77,fizbuz,79,80,fizbuz,82,83,fizbuz,85,86,fizbuz,88,89,fizbuz,91,92,fizbuz,94,95,fizbuz,97,98,fizbuz,100,"
	actual := fizzbuzz(100, options)
	if actual != expected {
		t.Errorf("Expected: %q\ngot: %q", expected, actual)
	}
}

func TestNegativeLimit(t *testing.T) {
	options := FizzbuzzOptions{
		Int1: 3,
		Int2: 3,
		Str1: "fiz",
		Str2: "buz",
	}
	expected := ""
	actual := fizzbuzz(-100, options)
	if actual != expected {
		t.Errorf("Expected: %q\ngot: %q", expected, actual)
	}
}

func TestZeroValues(t *testing.T) {
	defer func() {
		if r := recover(); r != "int1 and int2 cannot be 0" {
			t.Errorf("The code did not panic correctly")
		}
	}()
	options := FizzbuzzOptions{
		Int1: 3,
		Int2: 0,
		Str1: "fiz",
		Str2: "buz",
	}
	fizzbuzz(100, options)
}
