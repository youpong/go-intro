package intro_test

import (
	"testing"

	"github.com/youpong/go-intro"
)

func TestTruncateAtFinalSlash(t *testing.T) {
	pathName := intro.Path("/usr/bin/tso")
	pathName.TruncateAtFinalSlash()
	if !pathName.Equal(intro.Path("/usr/bin")) {
		t.Errorf("")
	}
}

func TestToUpper(t *testing.T) {
	pathName := intro.Path("/usr/bin/tso")
	pathName.ToUpper()
	if !pathName.Equal(intro.Path("/USR/BIN/TSO")) {
		t.Errorf("")
	}
}
