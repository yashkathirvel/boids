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
