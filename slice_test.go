package main_test

import (
	"bytes"
	"testing"
)

type path []byte

func (p path) Equal(other path) bool {
	return bytes.Equal(p, other)
}

func TestTruncateAtFinalSlash(t *testing.T) {
	pathName := path("/usr/bin/tso")
	pathName.TruncateAtFinalSlash()
	if !pathName.Equal(path("/usr/bin")) {
		t.Errorf("")
	}
}

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}
