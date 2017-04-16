package grid

import (
	"testing"
)

func TestMode(t *testing.T) {
	var o Mode
	e := modes(o)
	if len(e) != 0 {
		t.Errorf("expected 0 got %d", len(e))
	}
	e = modes(Default)
	if len(e) != 1 {
		t.Errorf("expected 1 got %d", len(e))
	}
	if e[0] != Default.String() {
		t.Errorf("expected empty string got %s", e[0])
	}
	e = modes(Default | Desktop)
	if e[1] != Desktop.String() {
		t.Errorf("expected %s got %s", e[1])
	}
}
