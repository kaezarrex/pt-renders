package main

import "github.com/fogleman/pt/pt"

// http://www.turbosquid.com/3d-models/free-brandy-cognac-glass-3d-model/741144

func main() {
	scene := pt.Scene{}
	material := pt.TransparentMaterial(pt.Color{}, 2, pt.Radians(0), 0.5)
	mesh, _ := pt.LoadOBJ("models/glass/brandy.obj", material)
	mesh.FitInside(pt.Box{pt.Vector{-1, -0.5, -1}, pt.Vector{1, 1, 1}}, pt.Vector{0.5, 0.5, 0.5})
	mesh.SmoothNormals()
	scene.Add(mesh)
	scene.Add(pt.NewSphere(pt.Vector{-2, 10, -5}, 1, pt.LightMaterial(pt.Color{1, 1, 1}, 12, pt.QuadraticAttenuation(0.1))))
	scene.Add(pt.NewCube(pt.Vector{-1000, -1, -100}, pt.Vector{100, -0.5, 100}, pt.DiffuseMaterial(pt.Color{1, 1, 1})))
	camera := pt.LookAt(pt.Vector{0, 1, -3}, pt.Vector{0, 0, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("renders/brandy%03d.png", 1000, &scene, &camera, 1024, 1024, -1, 16, 8)
}
