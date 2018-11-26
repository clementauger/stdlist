package stdlist_test

import (
	"testing"

	"github.com/clementauger/stdlist"
)

func TestHas(t *testing.T) {

	if !stdlist.Has("go1.1", "bytes") {
		t.Fatal("invalid result")
	}

	if !stdlist.Has("1.1", "bytes") {
		t.Fatal("invalid result")
	}

	if stdlist.Has("1.1", "NoSuchPackage") {
		t.Fatal("invalid result")
	}

}
