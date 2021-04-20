package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

var (
	alphabetLength = 26   // English alphabet length
	minWeight      = 0.01 // Min Weight Value
	maxWeight      = 1.0  // Max Weight Value
	alfa           = 0.01 // Learing rate
)

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func GenerateWeights() []float64 {
	// fmt.Println(randFloats(minWeight, maxWeight, alphabetLength))
	return randFloats(minWeight, maxWeight, alphabetLength)
}

func DotProduct(lett []float64, per *Perceptron) float64 {
	var sum float64
	if len(lett) != len(per.Weight) {
		fmt.Println("Can not apply dot product. Different dimensions.")
		os.Exit(1)
	}
	if lett == nil || per.VectorIsNull() {
		fmt.Println("Empty vector(s).")
		os.Exit(1)
	}

	for i := 0; i < len(per.Weight); i++ {
		sum += lett[i] * per.Weight[i]
	}
	return sum
}

func (pr *Perceptron) Normalize() {
	var weightSum float64
	for i := 0; i < len(pr.Weight); i++ {
		weightSum += pr.Weight[i]
	}
	normilizeCoef := 1.0 / weightSum
	// fmt.Println(normilizeCoef)
	// var val float64
	for j := 0; j < len(pr.Weight); j++ {
		pr.Weight[j] = pr.Weight[j] * normilizeCoef
		// val += pr.Weight[j]
		// fmt.Println(val)
	}
}

func (per *Perceptron) Predict(lett []float64) float64 {
	netVal := DotProduct(lett, per) - per.Treashold
	// fmt.Println(1.0 / (1.0 + math.Exp(-netVal)))
	return 1.0 / (1.0 + math.Exp(-netVal))
}

func (per *Perceptron) DeltaRule(letterProp []float64, desired float64, output float64) {
	// fmt.Println(per.Weight)
	// per.Normalize()
	// fmt.Println(per.Weight)
	err := desired - output
	for i := range per.Weight {
		per.Weight[i] += err * alfa * letterProp[i]
	}
	// fmt.Println()
	per.Treashold += err * alfa * (-1.0)
	per.Normalize()
}
