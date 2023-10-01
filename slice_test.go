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

func TestInsert(t *testing.T) {
	slice := make([]int, 5, 10)
	for i := range slice {
		slice[i] = i
	}
	slice = intro.Insert(slice, 2, 99)

	for i, e := range []int{0, 1, 99, 2, 3, 4} {
		if slice[i] != e {
			t.Errorf("slice[%d] == %d, but want %d\n", i, slice[i], e)
		}
	}
}
