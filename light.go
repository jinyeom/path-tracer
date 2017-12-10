package main

type DirectionalLight struct {
	color     *Vec3
	direction *Vec3
}

type PointLight struct {
	DirectionalLight
	position *Vec3
}
