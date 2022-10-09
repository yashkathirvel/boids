package main

import (
	"math"
	"reflect"
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		p1 OrderedPair
		p2 OrderedPair
	}

	type test struct {
		name string
		args args
		want float64
	}

	tests := make([]test, 2, 10)
	// TODO: Add test cases.

	// test0
	tests[0].name = "test0"
	tests[0].args.p1 = OrderedPair{1.0, 1.0}
	tests[0].args.p2 = OrderedPair{1.0, 1.0}
	tests[0].want = 0.0

	// test1
	tests[1].name = "test1"
	tests[1].args.p1 = OrderedPair{1.0, 5.0}
	tests[1].args.p2 = OrderedPair{5.0, 5.0}
	tests[1].want = 4.0

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestCopySky(t *testing.T) {
	type args struct {
		currentSky Sky
	}
	tests := []struct {
		name string
		args args
		want Sky
	}{
		// TODO: Add test cases.
		// Should return the same Sky
		{ // test0
			name: "test0",
			args: args{
				currentSky: Sky{
					width:            1.0,
					maxBoidSpeed:     1.0,
					proximity:        1.0,
					separationFactor: 1.0,
					alignmentFactor:  1.0,
					cohesionFactor:   1.0,
					boids: []Boid{
						{
							position:     OrderedPair{1.0, 2.0},
							velocity:     OrderedPair{1.0, 2.0},
							acceleration: OrderedPair{1.0, 2.0},
						},
						{
							position:     OrderedPair{1.0, 2.0},
							velocity:     OrderedPair{1.0, 2.0},
							acceleration: OrderedPair{1.0, 2.0},
						},
					},
				},
			},

			want: Sky{
				width:            1.0,
				maxBoidSpeed:     1.0,
				proximity:        1.0,
				separationFactor: 1.0,
				alignmentFactor:  1.0,
				cohesionFactor:   1.0,
				boids: []Boid{
					{
						position:     OrderedPair{1.0, 2.0},
						velocity:     OrderedPair{1.0, 2.0},
						acceleration: OrderedPair{1.0, 2.0},
					},
					{
						position:     OrderedPair{1.0, 2.0},
						velocity:     OrderedPair{1.0, 2.0},
						acceleration: OrderedPair{1.0, 2.0},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopySky(tt.args.currentSky); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopySky() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateVelocity(t *testing.T) {
	type args struct {
		b            Boid
		time         float64
		maxBoidSpeed float64
	}
	tests := []struct {
		name string
		args args
		want OrderedPair
	}{
		// TODO: Add test cases.
		{ // test0
			name: "test0",
			args: args{
				b: Boid{
					position:     OrderedPair{2.5, 5.0},
					velocity:     OrderedPair{-1.0, 2.0},
					acceleration: OrderedPair{1.0, 2.0},
				},
				time:         1.0,
				maxBoidSpeed: 2.0,
			},
			want: OrderedPair{0.0 * 2.0 / math.Sqrt(0.0*0.0+4.0*4.0), 4.0 * 2.0 / math.Sqrt(0.0*0.0+4.0*4.0)},
		},

		{ // test1
			name: "test1",
			args: args{
				b: Boid{
					position:     OrderedPair{2.5, 5.0},
					velocity:     OrderedPair{1.0, 1.0},
					acceleration: OrderedPair{1.0, 1.0},
				},
				time:         1.0,
				maxBoidSpeed: 2.0,
			},
			want: OrderedPair{2.0 * 2.0 / math.Sqrt(8), 2.0 * 2.0 / math.Sqrt(8)},
		},

		{ // test2
			name: "test2",
			args: args{
				b: Boid{
					position:     OrderedPair{2.5, 5.0},
					velocity:     OrderedPair{1.0, 1.0},
					acceleration: OrderedPair{0.2, 0.2},
				},
				time:         1.0,
				maxBoidSpeed: 2.0,
			},
			want: OrderedPair{1.2, 1.2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateVelocity(tt.args.b, tt.args.time, tt.args.maxBoidSpeed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateVelocity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdatePosition(t *testing.T) {
	type args struct {
		b        Boid
		time     float64
		skyWidth float64
	}
	tests := []struct {
		name string
		args args
		want OrderedPair
	}{
		// TODO: Add test cases.
		{ // test0
			name: "test0",
			args: args{
				b: Boid{
					position:     OrderedPair{1.0, 2.0},
					velocity:     OrderedPair{1.0, 2.0},
					acceleration: OrderedPair{1.0, 2.0},
				},
				time:     1.0,
				skyWidth: 20.0,
			},
			want: OrderedPair{2.5, 5.0},
		},
		{ // test1
			name: "test1",
			args: args{
				b: Boid{
					position:     OrderedPair{19.0, 1.0},
					velocity:     OrderedPair{1.0, 1.0},
					acceleration: OrderedPair{1.0, 1.0},
				},
				time:     1.0,
				skyWidth: 20.0,
			},
			want: OrderedPair{0.5, 2.5},
		},
		{ // test2
			name: "test2",
			args: args{
				b: Boid{
					position:     OrderedPair{1.0, 19.0},
					velocity:     OrderedPair{1.0, 1.0},
					acceleration: OrderedPair{1.0, 1.0},
				},
				time:     1.0,
				skyWidth: 20.0,
			},
			want: OrderedPair{2.5, 0.5},
		},
		{ // test3
			name: "test3",
			args: args{
				b: Boid{
					position:     OrderedPair{19.0, 19.0},
					velocity:     OrderedPair{1.0, 1.0},
					acceleration: OrderedPair{1.0, 1.0},
				},
				time:     1.0,
				skyWidth: 20.0,
			},
			want: OrderedPair{0.5, 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdatePosition(tt.args.b, tt.args.time, tt.args.skyWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args struct {
		inputs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				inputs: []float64{1.0, 2.0},
			},
			want: 1.5,
		},
		{
			name: "test1",
			args: args{
				inputs: []float64{1.0},
			},
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.inputs); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeNetForce(t *testing.T) {
	type args struct {
		currentSky Sky
		b          Boid
	}
	tests := []struct {
		name string
		args args
		want OrderedPair
	}{
		// TODO: Add test cases.
		{
			name: "test0", // no forces
			args: args{
				currentSky: Sky{
					width: 500.0,
					boids: []Boid{
						{position: OrderedPair{100.0, 400.0}, velocity: OrderedPair{12.0, -16.0}, acceleration: OrderedPair{0.0, 0.0}},
						{position: OrderedPair{350.0, 150.0}, velocity: OrderedPair{-30.0, 0.0}, acceleration: OrderedPair{0.0, 0.0}},
					},
					maxBoidSpeed:     10.0,
					proximity:        100.0,
					separationFactor: 1.0,
					alignmentFactor:  1.0,
					cohesionFactor:   1.0,
				},

				b: Boid{
					position:     OrderedPair{19.0, 1.0},
					velocity:     OrderedPair{12.0, -16.0},
					acceleration: OrderedPair{0.0, 0.0},
				},
			},
			want: OrderedPair{0.0, 0.0},
		},
		{
			name: "test1", // forces present
			args: args{
				currentSky: Sky{
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
				},

				b: Boid{position: OrderedPair{241.864115, 376.203635}, velocity: OrderedPair{-0.511419, -0.859332}, acceleration: OrderedPair{0.0, 0.0}},
			},
			want: OrderedPair{-0.5913, -0.0514},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputeNetForce(tt.args.currentSky, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeNetForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
