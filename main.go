package main

// func main() {

// 	// read command line arguments in the following order:
// 	// numBoids skyWidth initialSpeed maxBoidSpeed numGens proximity separationFactor
// 	// alignmentFactor cohesionFactor canvasWidth imageFrequency

// 	numBoids, err1 := strconv.Atoi(os.Args[1])
// 	if err1 != nil {
// 		panic(err1)
// 	}
// 	if numBoids < 0 {
// 		panic("Negative number of boids given.")
// 	}

// 	//os.Args[2] is going to be time step parameter
// 	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
// 	if err2 != nil {
// 		panic(err2)
// 	}

// 	//os.Args[3] is going to be the initial speed of the Boids
// 	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
// 	if err3 != nil {
// 		panic(err3)
// 	}

// 	//os.Args[4] is going to be the maximum speed of the Boids
// 	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
// 	if err4 != nil {
// 		panic(err4)
// 	}

// 	// os.Args[5] is going to be the number of generations the
// 	// simulation is run
// 	numGens, err5 := strconv.Atoi(os.Args[5])
// 	if err5 != nil {
// 		panic(err5)
// 	}
// 	if numGens < 0 {
// 		panic("Negative number of generations given.")
// 	}

// 	//os.Args[6] is going to be the proximity
// 	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
// 	if err6 != nil {
// 		panic(err6)
// 	}

// 	//os.Args[7] is going to be separation factor
// 	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
// 	if err7 != nil {
// 		panic(err7)
// 	}
// 	//os.Args[8] is going to be slignment factor
// 	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
// 	if err8 != nil {
// 		panic(err8)
// 	}

// 	//os.Args[9] is going to be cohesion factor
// 	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
// 	if err9 != nil {
// 		panic(err9)
// 	}

// 	//os.Args[10] is the canvas width
// 	canvasWidth, err10 := strconv.Atoi(os.Args[10])
// 	if err10 != nil {
// 		panic(err10)
// 	}

// 	//os.Args[11] is how often to make a canvas
// 	imageFrequency, err11 := strconv.Atoi(os.Args[11])
// 	if err11 != nil {
// 		panic(err11)
// 	}

// 	// temporary
// 	// fmt.Println(numBoids, skyWidth, initialSpeed, maxBoidSpeed, numGens, proximity, separationFactor, alignmentFactor, cohesionFactor, canvasWidth, imageFrequency)

// 	// Start the simulation
// 	fmt.Println("Command line arguments read successfully.")

// 	// Create initialSky
// 	system := CreateSky(skyWidth, numBoids, initialSpeed, maxBoidSpeed, proximity, separationFactor, alignmentFactor, cohesionFactor)

// 	fmt.Println("Simulating system.")

// 	timePoints := SimulateBoids(system, numGens, 0.1)

// 	fmt.Println("Sky has been simulated!")
// 	fmt.Println("Ready to draw images.")

// 	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

// 	fmt.Println("Images drawn!")

// 	fmt.Println("Making GIF.")

// 	gifhelper.ImagesToGIF(images, "four")

// 	fmt.Println("Animated GIF produced!")

// 	fmt.Println("Exiting normally.")
// }

func main() {
	tcnf()
}
