package model

import "sync"

type calculater func([]float64) (float64, error)

//Layer represents a collection of neurons (Cells) in a given domain of a
//network
type Layer struct {
	neurons    []Neuron
	inputSize  int
	outputSize int
}

//CalculateOutput gets the output array for a layer base on individual Neurons
//TODO: Use concurrent calculations
func (l *Layer) CalculateOutput(inputs []float64) ([]float64, error) {
	output := make([]float64, l.outputSize)
	var wg sync.WaitGroup
	wg.Add(l.outputSize)
	for index, neuron := range l.neurons {
		go func(n Neuron, i int) {
			defer wg.Done()
			readOutputToChan(n.Process, inputs, output, i)
		}(neuron, index)
	}

	wg.Wait()
	return output, nil
}

func readOutputToChan(
	f calculater, input []float64, output []float64, index int) {
	value, err := f(input)
	if err != nil {
		panic(err)
	}
	output[index] = value
}

//Breed takes two parent layers and attempts to breed children
func (l *Layer) Breed(pairedLayer Layer) error {
	if l.inputSize != pairedLayer.inputSize {
		return NewMismatchError(l.inputSize, pairedLayer.inputSize)
	}
	if l.outputSize != pairedLayer.outputSize {
		return NewMismatchError(l.outputSize, pairedLayer.inputSize)
	}

	neurons := make([]Neuron, l.outputSize)
	var wg sync.WaitGroup
	wg.Add(l.outputSize)
	for index, neuron := range l.neurons {
		go func(n1 Neuron, n2 Neuron, index int) {
			defer wg.Done()
			child, err := n1.Breed(n2)
			if err != nil {
				panic(err)
			}
			neurons[index] = child
		}(neuron, pairedLayer.neurons[index], index)
	}

	wg.Wait()
	l.neurons = neurons
	return nil
}

//NewLayer creates a new layer of neurons
func NewLayer(
	size int, sizeOfPreviousLayer int, cellCharacter *CellCharacter) Layer {

	neurons := make([]Neuron, 0, size)
	for i := 0; i < size; i++ {
		neurons = append(
			neurons, *NewCell(sizeOfPreviousLayer, cellCharacter))
	}
	return Layer{neurons, sizeOfPreviousLayer, size}
}
