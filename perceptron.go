package main

type Perceptron struct {
	Label     string
	Treashold float64
	Weight    []float64
}

func (per *Perceptron) VectorIsNull() bool {
	return per.Weight == nil
}
