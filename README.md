# A Genetic Algorithm Optimized, Neural Network Processor
This library contains the root components for a neural network to be created and optimized.

# High level usage
```GO
  package main

  import (
    "fmt"
    "math"
    "math/rand"

    m "github.com/JamesRCooper/geneticNeuralNetwork/model"
  )

  func main() {
    func main() {
    	inputs := []float64{0.57735, 0.57735, 0.57735}
    	desiredOutput := []float64{0.45882, 0.10196, 0.69804}

    	layerSizes := []int{4, 3}
    	cellCharacteristics := m.CellCharacter{
    		MutationRate: 0.015,
    		Activater:    sigmoid,
    		GeneCreator:  geneCreator,
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

    func sigmoid(operand float64) float64 {
    	return 1.0 / (1.0 + math.Exp(operand))
    }

    func geneCreator() float64 {
    	return 2*rand.Float64() - 1
    }
  }
```
