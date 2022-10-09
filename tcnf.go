package main

import "fmt"

func tcnf() {

	currentSky := Sky{
		width: 400.0,
		boids: []Boid{
			{position: OrderedPair{241.864115, 376.203635}, velocity: OrderedPair{-0.511419, -0.859332}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{175.085675, 169.854999}, velocity: OrderedPair{-0.386609, -0.922244}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{26.254808, 62.607702}, velocity: OrderedPair{0.820062, 0.572275}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{120.364744, 206.085051}, velocity: OrderedPair{0.389291, -0.921115}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{85.705549, 152.262876}, velocity: OrderedPair{-0.414708, 0.909954}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{187.555938, 113.21366}, velocity: OrderedPair{-0.267519, 0.963553}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{271.63387, 87.421221}, velocity: OrderedPair{0.289913, 0.957053}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{144.348567, 228.26931}, velocity: OrderedPair{0.649407, -0.760441}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{117.245698, 118.833025}, velocity: OrderedPair{0.016166, -0.999869}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{82.633065, 346.134005}, velocity: OrderedPair{-0.328555, -0.944485}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{209.528122, 11.321233}, velocity: OrderedPair{0.544666, 0.838653}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{242.901376, 390.096648}, velocity: OrderedPair{0.877955, 0.478742}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{237.923439, 23.648261}, velocity: OrderedPair{-0.356268, -0.934384}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{120.609072, 69.306495}, velocity: OrderedPair{-0.966841, -0.255377}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{217.662229, 111.403049}, velocity: OrderedPair{-0.885676, 0.464304}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{212.234286, 101.4162}, velocity: OrderedPair{-0.200209, 0.979753}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{315.441966, 144.722192}, velocity: OrderedPair{0.7313, -0.682055}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{118.844904, 357.744692}, velocity: OrderedPair{0.818314, 0.574772}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{390.766747, 29.7164}, velocity: OrderedPair{0.173232, 0.984881}, acceleration: OrderedPair{0.0, 0.0}},
			{position: OrderedPair{272.431325, 96.606035}, velocity: OrderedPair{-0.377002, 0.926213}, acceleration: OrderedPair{0.0, 0.0}},
		},
		maxBoidSpeed:     2.0,
		proximity:        200.0,
		separationFactor: 1.0,
		alignmentFactor:  1.0,
		cohesionFactor:   1.0,
	}

	b := Boid{position: OrderedPair{241.864115, 376.203635}, velocity: OrderedPair{-0.511419, -0.859332}, acceleration: OrderedPair{0.0, 0.0}}

	fmt.Println(ComputeNetForce(currentSky, b))

}
