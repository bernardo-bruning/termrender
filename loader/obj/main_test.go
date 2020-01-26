package obj_test

import (
	"github.com/bernardo-bruning/termrender/loader/obj"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	file, err := os.Open("cube.obj")
	if err != nil {
		t.Error(err)
	}

	_, err = obj.Load(file)
	if err != nil {
		t.Error(err)
	}
}
