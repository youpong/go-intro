package main_test

import (
	"bytes"
	"testing"
)

type path []byte

func TestTruncateAtFinalSlash(t *testing.T) {
	pathName := path("/usr/bin/tso")
	pathName.TruncateAtFinalSlash()
	if !pathName.Equal(path("/usr/bin")) {
		t.Errorf("")
	}
}

func TestToUpper(t *testing.T) {
	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	if !pathName.Equal(path("/USR/BIN/TSO")) {
		t.Errorf("")
	}
}

func (p path) Equal(other path) bool {
	return bytes.Equal(p, other)
}

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p *path) ToUpper() {
	for i, e := range *p {
		if 'a' <= e && e <= 'z' {
			(*p)[i] = e + 'A' - 'a'
		}
	}
}
