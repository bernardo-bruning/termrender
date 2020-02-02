package obj

import (
	"bufio"
	"errors"
	"github.com/bernardo-bruning/termrender/render"
	"io"
	"strconv"
	"strings"
)

func loadVector(values []string) (render.Vector, error) {
	if len(values) < 3 {
		return render.Vector{}, errors.New("Invalid load vector")
	}

	x, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return render.Vector{}, err
	}

	y, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return render.Vector{}, err
	}

	if len(values) == 3 {
		return render.Vector{X: x, Y: y, Z: 0}, nil
	}

	z, err := strconv.ParseFloat(values[3], 64)
	if err != nil {
		return render.Vector{}, err
	}

	return render.Vector{X: x, Y: y, Z: z}, nil
}

func loadTriangle(vectors []render.Vector, a, b, c string) (render.Triangle, error) {
	ai, err := strconv.Atoi(a)
	if err != nil {
		return render.Triangle{}, err
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		return render.Triangle{}, err
	}

	ci, err := strconv.Atoi(c)
	if err != nil {
		return render.Triangle{}, err
	}

	return render.NewTriangle(
		vectors[ai-1],
		vectors[bi-1],
		vectors[ci-1],
	), nil
}

//Load return a mesh from file obj
func Load(r io.Reader) (render.Mesh, error) {
	scanner := bufio.NewScanner(r)
	vectors := []render.Vector{}
	vectorsTexture := []render.Vector{}
	triangles := []render.Triangle{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		obj := strings.Split(line, " ")
		if strings.EqualFold(obj[0], "v") {
			vector, err := loadVector(obj)
			if err != nil {
				return render.Mesh{}, err
			}
			vectors = append(vectors, vector)
		}

		if strings.EqualFold(obj[0], "vt") {
			vector, err := loadVector(obj)
			if err != nil {
				return render.Mesh{}, err
			}

			vectorsTexture = append(vectorsTexture, vector)
		}

		if strings.EqualFold(obj[0], "f") {
			if len(obj) < 4 {
				return render.Mesh{}, errors.New("Invalid load face")
			}
			a := strings.Split(obj[1], "/")
			b := strings.Split(obj[2], "/")
			c := strings.Split(obj[3], "/")
			
			triangle, err := loadTriangle(vectors, a[0], b[0], c[0])
			if err != nil {
				return render.Mesh{}, nil
			}

			triangles = append(triangles, triangle)
		}
	}

	return render.NewMesh(triangles), nil
}
