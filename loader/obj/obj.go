package obj

import (
	"bufio"
	"errors"
	"github.com/bernardo-bruning/termrender/render"
	"io"
	"strconv"
	"strings"
)

//Load return a mesh from file obj
func Load(r io.Reader) (render.Mesh, error) {
	scanner := bufio.NewScanner(r)
	vectors := []render.Vector{}
	triangles := []render.Triangle{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		obj := strings.Split(line, " ")
		if strings.EqualFold(obj[0], "v") {
			if len(obj) < 4 {
				return render.Mesh{}, errors.New("Invalid load vertice")
			}

			x, err := strconv.ParseFloat(obj[1], 64)
			if err != nil {
				return render.Mesh{}, err
			}

			y, err := strconv.ParseFloat(obj[2], 64)
			if err != nil {
				return render.Mesh{}, err
			}

			z, err := strconv.ParseFloat(obj[3], 64)
			if err != nil {
				return render.Mesh{}, err
			}

			vectors = append(vectors, render.Vector{x, y, z})
		}

		if strings.EqualFold(obj[0], "f") {
			if len(obj) < 4 {
				return render.Mesh{}, errors.New("Invalid load face")
			}
			a := strings.Split(obj[1], "/")
			b := strings.Split(obj[2], "/")
			c := strings.Split(obj[3], "/")

			ai, err := strconv.Atoi(a[0])
			if err != nil {
				return render.Mesh{}, err
			}

			bi, err := strconv.Atoi(b[0])
			if err != nil {
				return render.Mesh{}, err
			}

			ci, err := strconv.Atoi(c[0])
			if err != nil {
				return render.Mesh{}, err
			}

			triangles = append(triangles, render.NewTriangle(
				vectors[ai-1],
				vectors[bi-1],
				vectors[ci-1],
			))
		}
	}

	return render.NewMesh(triangles), nil
}
