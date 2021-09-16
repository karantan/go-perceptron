package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ISIZE = 2
const WSIZE = ISIZE + 1 // weights + bias
const LEARNING_RATE = 0.1
const ITERATIONS = 10

var weights [WSIZE]float64

func main() {

	fmt.Printf("Starting weights %g %g bias %g\n", weights[0], weights[1], weights[2])
	train()
	fmt.Printf("Final weights %g %g bias %g\n", weights[0], weights[1], weights[2])
}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < WSIZE; i++ {
		weights[i] = rand.Float64()
	}
}

func feedforward(inputs []int) int {
	sum := 0.0

	//Calculate inputs * weights
	for i := 0; i < ISIZE; i++ {
		sum += weights[i] * float64(inputs[i])
	}
	// Add in the bias - last element in weights
	sum += weights[WSIZE-1]

	// Actiovation function. 1 if sum >= 1
	if sum >= 1.0 {
		return 1
	}
	return 0
}

func train() {

	var (
		iterations      int
		iteration_error int
		desired_output  int
		output          int
		err             int
	)

	// train the boolean OR set
	test := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for {
		iteration_error = 0 // reset the iteration error
		fmt.Printf("Iteration %d\n", iterations)

		for i := 0; i < len(test); i++ {
			if test[i][0] == 1 || test[i][1] == 1 {
				desired_output = 1
			} else {
				desired_output = 0
			}
			output = feedforward(test[i])

			err = desired_output - output

			fmt.Printf("%d or %d = %d (desired output: %d)\n", test[i][0], test[i][1], output, desired_output)

			weights[0] += (LEARNING_RATE * float64(err) * float64(test[i][0]))
			weights[1] += (LEARNING_RATE * float64(err) * float64(test[i][1]))
			weights[2] += (LEARNING_RATE * float64(err))
			fmt.Printf("Adjusted weights %g %g bias %g\n", weights[0], weights[1], weights[2])

			iteration_error += err
		}

		fmt.Printf("Iteration error %d\n\n", iteration_error)
		iterations++

		// max `ITERATIONS` loops even if we still have error
		if iteration_error == 0 || iterations >= ITERATIONS {
			break
		}
	}
}
