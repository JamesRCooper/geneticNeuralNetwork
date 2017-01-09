package model

import "math/rand"

type (
	//Breeder is used to denote a class that can breed with any other class
	//containing the same number of genes
	Breeder interface {
		Breed(partner Breeder) (Neuron, error)
		NumberOfGenes() int
		NthGene(n int) float64
	}

	//Processor is any class that can generate an ouput given an input array
	Processor interface {
		Process(inputs []float64) (float64, error)
		NumberOfInputs() int
	}

	//Neuron is a any class that can process inputs, and breed
	Neuron interface {
		Breeder
		Processor
	}
)

//Cell represents an inner neuron, where weights are the weights placed on the
//inputs
type Cell struct {
	Weights   []float64
	Size      int
	character *CellCharacter
}

//Process gets the output value of the Cell given some set of inputs
func (c Cell) Process(inputs []float64) (float64, error) {
	numOfInputs := len(inputs)
	if c.Size != numOfInputs {
		return 0, NewMismatchError(c.Size, numOfInputs)
	}
	var sum float64
	for index, weight := range c.Weights {
		sum += weight * inputs[index]
	}
	return c.character.Activater(sum), nil
}

//NumberOfInputs returns the number of inputs a cell can receive
func (c Cell) NumberOfInputs() int {
	return c.Size
}

//Breed creates a new cell by mixing the genes of this cell and a pairedCell,
//along with mutations based on the MutationRate
func (c Cell) Breed(pairedCell Breeder) (Neuron, error) {
	if c.Size != pairedCell.NumberOfGenes() {
		return *new(Neuron), NewMismatchError(c.Size, pairedCell.NumberOfGenes())
	}

	genes := make([]float64, c.Size)
	for index := 0; index < c.Size; index++ {
		cellChoice := rand.Float64()
		if cellChoice < c.character.MutationRate {
			genes[index] = c.character.GeneCreator()
		} else if cellChoice < (c.character.MutationRate/2.0)+0.5 {
			genes[index] = c.Weights[index]
		} else {
			genes[index] = pairedCell.NthGene(index)
		}
	}

	return Cell{genes, c.Size, c.character}, nil
}

//NumberOfGenes returns the number of the inputs on a cell
func (c Cell) NumberOfGenes() int {
	return c.Size
}

//NthGene returns the nth weight in a cell
func (c Cell) NthGene(n int) float64 {
	return c.Weights[n]
}

//NewCell creates a new instance of Cell
func NewCell(size int, character *CellCharacter) *Cell {

	newWeights := make([]float64, size)
	for i := range newWeights {
		newWeights[i] = character.GeneCreator()
	}
	return &Cell{newWeights, size, character}
}
