package main

import (
	"math"
	"math/rand"
)

//place your non-drawing functions here.

//SimulateBoids simulates the sky over a series of snap shots separated by equal unit time.
//Input: an initial Sky, a number of generations, and a time parameter (in seconds).
//Output: a slice of Sky objects

func SimulateBoids(initialSky Sky, numGens int, time float64) []Sky {

	// init Sky array with first element being the initial Sky
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialSky

	// now range over the number of generations and update the Sky each time
	for i := 1; i < numGens+1; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], time)
	}

	return timePoints
}

// UpdateSky
// Input: an initial Sky, and the time step (in seconds).
// Output: the Sky object simulating the boids at the current generation.

func UpdateSky(initialSky Sky, time float64) Sky {

	// init new Sky object and set it equal to the initial Sky object
	// (this statement has to be updated with a dedicated function to
	// copy the Sky object)
	newSky := CopySky(initialSky)
	// now range over the Boids in the Sky object and update their fields
	for i := range newSky.boids {
		newSky.boids[i].acceleration = UpdateAcceleration(initialSky, newSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(newSky.boids[i], time, initialSky.maxBoidSpeed)
		newSky.boids[i].position = UpdatePosition(newSky.boids[i], time, initialSky.width)
	}
	return newSky
}

//UpdateAcceleration
//Input: Sky object and a Boid B
//Output: The net acceleration on B due to interactions with other Boids

func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(currentSky, b)

	//now, calculate acceleration (F = ma)
	// here, m = 1
	accel.x = force.x
	accel.y = force.y

	return accel
}

// ComputeNetForce
// Input: A slice of Boid objects and proximityFactor from current Sky object
// and an individual Boid
// Output: The net force vector (OrderedPair) acting on the given Boid
// due to separation, alignment, and cohesion from other Boids if within
// proximityFactor

func ComputeNetForce(currentSky Sky, b Boid) OrderedPair {
	numBoidsAffected := 0
	var netForce OrderedPair

	for _, i := range currentSky.boids {
		dist := Distance(b.position, i.position)
		// fmt.Println(dist)
		//only do a force computation if current boid is not the input Boid
		if i != b && dist < currentSky.proximity {
			// compute the three rules
			sepForce := SeparationForce(b, i, currentSky.separationFactor)
			alnForce := AlignmentForce(b, i, currentSky.alignmentFactor)
			cohForce := CohesionForce(b, i, currentSky.cohesionFactor)
			// add to netForce
			netForce.x += sepForce.x + alnForce.x + cohForce.x
			netForce.y += sepForce.y + alnForce.y + cohForce.y
			// add 1 to boids counter
			numBoidsAffected++
			// fmt.Println(i, numBoidsAffected, dist, separationForce, alignmentForce, cohesionForce)
		}
	}
	if numBoidsAffected != 0 {
		netForce.x /= float64(numBoidsAffected)
		netForce.y /= float64(numBoidsAffected)
	}
	return netForce
}

// SeparationForce
// Input:
// Output:

func SeparationForce(b1, b2 Boid, separationFactor float64) OrderedPair {
	var sepForce OrderedPair
	dist := Distance(b1.position, b2.position)

	sepForce.x = separationFactor * (b1.position.x - b2.position.x) / (dist * dist)
	sepForce.y = separationFactor * (b1.position.y - b2.position.y) / (dist * dist)

	return sepForce
}

// AlignmentForce
// Input:
// Output:

func AlignmentForce(b1, b2 Boid, alignmentFactor float64) OrderedPair {
	var alnForce OrderedPair
	dist := Distance(b1.position, b2.position)

	alnForce.x = alignmentFactor * b2.velocity.x / dist
	alnForce.y = alignmentFactor * b2.velocity.y / dist

	return alnForce

}

// CohesionForce
// Input:
// Output:

func CohesionForce(b1, b2 Boid, cohesionFactor float64) OrderedPair {

	var cohForce OrderedPair
	dist := Distance(b1.position, b2.position)

	cohForce.x = cohesionFactor * (b2.position.x - b1.position.x) / dist
	cohForce.y = cohesionFactor * (b2.position.y - b1.position.y) / dist

	return cohForce

}

// UpdateVelocity
// Input: a Boid object and a time step (float64).
// Output: The orderedPair corresponding to the velocity of this object
// after a single time step, using the Boid's current acceleration.

func UpdateVelocity(b Boid, time, maxBoidSpeed float64) OrderedPair {
	var vel OrderedPair

	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

	// check if velocity exceeds the maxBoidSpeed
	net := math.Sqrt(vel.x*vel.x + vel.y*vel.y)

	if net > maxBoidSpeed { // if velocity exceeds the maxBoidSpeed,
		// multiply components by (maxBoidSpeed/net)
		vel.x *= maxBoidSpeed / net
		vel.y *= maxBoidSpeed / net
	}

	return vel
}

// UpdatePosition
// Input: a Boid b and a time step (float64).
// Output: The OrderedPair corresponding to the updated position of the Boid after a
// single time step, using the boid's current acceleration and velocity.

func UpdatePosition(b Boid, time, skyWidth float64) OrderedPair {
	var pos OrderedPair
	// update position fields on Boid
	pos.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	pos.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y
	// check if position is out of bounds
	// this will simulate the torus
	if pos.x > skyWidth { // x component exceeds the sky width
		pos.x -= skyWidth
	}
	if pos.x < 0 { // x component is negative
		pos.x += skyWidth
	}
	if pos.y > skyWidth { // y component exceeds the sky width
		pos.y -= skyWidth
	}
	if pos.y < 0 { // y component is negative
		pos.y += skyWidth
	}
	return pos
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func Average(inputs []float64) float64 {
	var sum float64
	for _, i := range inputs {
		sum += i
	}
	return sum / float64(len(inputs))
}

// CreateSky
// Input: skyWidth, numBoids, initialSpeed, maxBoidSpeed,
// proximity, separationFactor, alignmentFactor, cohesionFactor
// Output: a new Sky object, with all the required fields set

func CreateSky(skyWidth float64, numBoids int, initialSpeed float64, maxBoidSpeed float64, proximity float64, separationFactor float64, alignmentFactor float64, cohesionFactor float64) Sky {
	var newSky Sky

	// set all the float64 fields
	newSky.width = skyWidth
	newSky.maxBoidSpeed = maxBoidSpeed
	newSky.proximity = proximity
	newSky.separationFactor = separationFactor
	newSky.alignmentFactor = alignmentFactor
	newSky.cohesionFactor = cohesionFactor

	// let's make the new Sky's slice of Boid objects
	newSky.boids = make([]Boid, numBoids)

	// now initialize all of the Boids
	for i := range newSky.boids {
		// init random positions
		randomPositionX := rand.Float64()*(skyWidth-0.0) + 0.0
		randomPositionY := rand.Float64()*(skyWidth-0.0) + 0.0
		// init random angle
		min := 0.0
		max := 2 * math.Pi
		randomAngle := rand.Float64()*(max-min) + min
		// fill in the fields
		newSky.boids[i].velocity.x = initialSpeed * math.Cos(randomAngle)
		newSky.boids[i].velocity.y = initialSpeed * math.Sin(randomAngle)
		newSky.boids[i].position.x = randomPositionX
		newSky.boids[i].position.y = randomPositionY
	}
	return newSky
}

func CopySky(currentSky Sky) Sky {
	var newSky Sky

	// copy all of the float64 fields into the new Sky
	newSky.width = currentSky.width
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed
	newSky.proximity = currentSky.proximity
	newSky.separationFactor = currentSky.separationFactor
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor

	//let's make the new Sky's slice of Boid objects
	numBoids := len(currentSky.boids)
	newSky.boids = make([]Boid, numBoids)

	//now, copy all of the boids' fields into our new boids
	for i := range currentSky.boids {
		newSky.boids[i].position.x = currentSky.boids[i].position.x
		newSky.boids[i].position.y = currentSky.boids[i].position.y
		newSky.boids[i].velocity.x = currentSky.boids[i].velocity.x
		newSky.boids[i].velocity.y = currentSky.boids[i].velocity.y
		newSky.boids[i].acceleration.x = currentSky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = currentSky.boids[i].acceleration.y
	}

	return newSky
}
