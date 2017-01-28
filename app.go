package main

import (
	"fmt"
	"math"

	m "github.com/JamesRCooper/geneticNeuralNetwork/model"
	"github.com/JamesRCooper/geneticNeuralNetwork/mutation"
)

func main() {
	inputs := []float64{0.57735, 0.57735, 0.57735}
	desiredOutput := []float64{0.45882, 0.10196, 0.69804}

	layerSizes := []int{4, 3}
	cellCharacteristics := m.CellCharacter{
		NeuronBreeder: mutation.BuildGaussianBreeder(
			0.01, 0.25, mutation.StandardGeneCreator),
		Activator:   mutation.LogisticSigmoidActivator,
		GeneCreator: mutation.StandardGeneCreator,
	}
	network := m.NewNetwork(3, layerSizes, &cellCharacteristics)
	value, err := network.CalculateOutput(inputs)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(value)
	fmt.Println(calculateOutputError(value, desiredOutput))
}

func calculateOutputError(
	actualOutputs []float64, expectedOutputs []float64) (float64, error) {
	if len(actualOutputs) != len(expectedOutputs) {
		return 0.0, m.NewMismatchError(len(actualOutputs), len(expectedOutputs))
	}
	var sum float64
	for index, value := range actualOutputs {
		sum += math.Abs(value - expectedOutputs[index])
	}
	return math.Sqrt(sum), nil
}
