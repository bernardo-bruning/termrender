package obj_test

import (
	"os"
	"testing"

	"github.com/bernardo-bruning/termrender/loader/obj"
)

func TestLoad(t *testing.T) {
	file, err := os.Open("cube.obj")
	if err != nil {
		t.Error(err)
	}

	mesh, err := obj.Load(file)
	if err != nil {
		t.Error(err)
	}

	if len(mesh.Triangles) != 12 {
		t.Error("Number of triangles is invalid!")
	}
}
