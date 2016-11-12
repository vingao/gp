package debugutil

import (
	"log"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	type Test struct {
		A int
		B bool
		C string
		D struct {
			aa int
			BB bool
		}
		E map[string]string
	}

	st := &Test{}

	buf := walk(st)
	log.Println("buf=%s\n", buf)
	if !strings.Contains(buf, "0: A int = 0") {
		t.Error("Wrong response from walk")
	}
	if !strings.Contains(buf, "B bool = false") {
		t.Error("Wrong response from walk")
	}
	if !strings.Contains(buf, "D struct { aa int; BB bool } = {0 false}") {
		t.Error("Wrong response from walk")
	}

}
