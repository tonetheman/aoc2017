package main

import (
	"testing"
)

// TestPart1 - tests part 1 in day1
func TestPart1(t *testing.T) {

	tt := func(s string, val int) {
		res := captcha(s)
		if res != val {
			t.Error("sum is not ", val, res)
		}
	}

	tt("1111", 4)
	tt("1122", 3)
	tt("1234", 0)
	tt("91212129", 9)

}

func TestPart2(t *testing.T) {

	if captcha2("1212") != 6 {
		t.Error("1212 is not 6")
	}
	if captcha2("1221") != 0 {
		t.Error("1221 is no 0")
	}
	if captcha2("123425") != 4 {
		t.Error("12345 is not 4")
	}
	if captcha2("123123") != 12 {
		t.Error("123123 is not 12")
	}
	if captcha2("12131415") != 4 {
		t.Error("12131415 is not 4")
	}

}
