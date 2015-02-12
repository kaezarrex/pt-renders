package main

import "github.com/fogleman/pt/pt"

func main() {
	scene := pt.Scene{}
	wall := pt.SpecularMaterial(pt.HexColor(0xFCFAE1), 2)
	floor := pt.DiffuseMaterial(pt.HexColor(0x334D5C))

	scene.Add(pt.NewSphere(pt.Vector{2500, 1500, 300}, 200, pt.LightMaterial(pt.Color{1, 1, 1}, 0.5, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{-2500, 1500, 300}, 200, pt.LightMaterial(pt.Color{1, 1, 1}, 0.5, pt.NoAttenuation)))
	scene.Add(pt.NewSphere(pt.Vector{0, 1500, 2500}, 200, pt.LightMaterial(pt.Color{1, 1, 1}, 0.5, pt.NoAttenuation)))

	scene.Add(pt.NewCube(pt.Vector{-3000, -1, -446}, pt.Vector{3000, 2000, -445}, wall))
	scene.Add(pt.NewCube(pt.Vector{-3000, -1, -446}, pt.Vector{3000, 0, 3000}, floor))
	scene.Add(pt.NewCube(pt.Vector{-3001, -1, -446}, pt.Vector{-3000, 2000, 3000}, wall))
	scene.Add(pt.NewCube(pt.Vector{3000, -1, -446}, pt.Vector{3001, 2000, 3000}, wall))
	scene.Add(pt.NewCube(pt.Vector{-3000, -1, 3000}, pt.Vector{3000, 2000, 30001}, wall))
	scene.Add(pt.NewCube(pt.Vector{-3000, 2000, -446}, pt.Vector{3000, 2001, 3000}, wall))

	//scene.Add(pt.NewCube(pt.Vector{-829, 0, -445}, pt.Vector{1881, 809, 445}, pt.SpecularMaterial(pt.HexColor(0x3AD83F), 2)))

	material := pt.GlossyMaterial(pt.Color{}, 1.5, pt.Radians(30))
	mesh, _ := pt.LoadOBJ("models/couch/couch.obj", material)
	mesh.SmoothNormals()
	mesh.Compile()
	scene.Add(mesh)

	camera := pt.LookAt(pt.Vector{1800, 1300, 1800}, pt.Vector{800, 300, 0}, pt.Vector{0, 1, 0}, 45)
	pt.IterativeRender("renders/couch%03d.png", 100, &scene, &camera, 2560/8, 1440/8, 0, 16, 4)
}
