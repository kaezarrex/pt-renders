package main

import (
	"fmt"

	"github.com/fogleman/pt/pt"
)

func main() {

	for i := 0; i < 90; i++ {

		scene := pt.Scene{}

		glass := pt.Material{pt.Color{}, nil, 0, pt.NoAttenuation, 2, pt.Radians(20), 0, true}

		cube := pt.NewCube(pt.Vector{-1, 0, -1}, pt.Vector{1, 2, 1}, glass)
		transform := pt.Rotate(pt.Vector{0, 1, 0}, pt.Radians(float64(i))).Translate(pt.Vector{-1, 0, 1})
		scene.Add(pt.NewTransformedShape(cube, transform))

		scene.Add(pt.NewSphere(pt.Vector{1.5, 1, 0}, 1, glass))
		scene.Add(pt.NewCube(pt.Vector{-100, -1, -100}, pt.Vector{100, 0, 100}, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
		scene.Add(pt.NewSphere(pt.Vector{-1, 4, -1}, 0.5, pt.LightMaterial(pt.Color{1, 1, 1}, 3, pt.LinearAttenuation(1))))
		camera := pt.LookAt(pt.Vector{0, 2, -5}, pt.Vector{0, 0, 3}, pt.Vector{0, 1, 0}, 45)
		pt.IterativeRender(fmt.Sprintf("out%03d.png", i), 1, &scene, &camera, 2560/8, 1440/8, 16, 64, 8)
	}
}
