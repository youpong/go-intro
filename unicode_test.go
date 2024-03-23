package intro_test

import (
	"testing"
	"unicode/utf8"
)

func TestXxx(t *testing.T) {
	s := "\u3042\U00020BB7\u30D5\u309A\U00020B9F"
	if len(s) != 17 {
		t.Errorf("%d", len(s))
	}
	if utf8.RuneCountInString(s) != 5 {
		t.Errorf("%d", utf8.RuneCountInString(s))
	}
}
