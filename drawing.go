package main

import (
	"canvas"
	"image"
)

//place your drawing code here.

// AnimateSystem takes a slice of Sky objects along with a canvas width
// parameter and a frequency parameter. It generates a slice of images corresponding to drawing each Universe whose index is divisible by the frequency parameter
// on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Sky, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { //only draw if current index of universe
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing
// a Sky object's bodies on a square canvas that is (canvasWidth.pixels)^2

func DrawToCanvas(s Sky, canvasWidth int) image.Image {

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the Boids and draw them
	for _, b := range s.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))
		cx := (b.position.x / s.width) * float64(canvasWidth)
		cy := (b.position.y / s.width) * float64(canvasWidth)
		// radius of boids is 5
		r := (5.0 / s.width) * float64(canvasWidth)

		// dimension of Boids
		c.Circle(cx, cy, r)
		c.Fill()
	}

	return c.GetImage()

}
