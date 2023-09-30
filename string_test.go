package main_test

import "testing"

func TestCompare(t *testing.T) {
	s := "ワカメ好き好き"
	if s == "お前はどこのワカメじゃ？" {
		t.Errorf("error")
	}

	var zero string
	if s == zero {
		t.Errorf("compare with zero")
	}

	if s == "" {
		t.Errorf("compare with empty")
	}
}

func TestCompre2(t *testing.T) {
	var zero string
	if zero == "" {
		t.Errorf("compare zero and empty")
	}
}
