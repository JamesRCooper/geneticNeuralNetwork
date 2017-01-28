package model

type (
	//Breeder is used to denote a class that can breed with any other class
	//containing the same number of genes
	Breeder interface {
		Breed(partner Breeder) (Neuron, error)
		Clone() Neuron
		NumberOfGenes() int
		NthGene(n int) float64
		CreateNew(genes []float64) Neuron
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
	weights   []float64
	size      int
	character *CellCharacter
}

//Process gets the output value of the Cell given some set of inputs
func (c Cell) Process(inputs []float64) (float64, error) {
	numOfInputs := len(inputs)
	if c.size != numOfInputs {
		return 0, NewMismatchError(c.size, numOfInputs)
	}
	var sum float64
	for index, weight := range c.weights {
		sum += weight * inputs[index]
	}
	return c.character.Activator(sum / float64(numOfInputs)), nil
}

//NumberOfInputs returns the number of inputs a cell can receive
func (c Cell) NumberOfInputs() int {
	return c.size
}

//CreateNew returns a new cell with the same character and size, but with the
//supplied set of genes
func (c Cell) CreateNew(genes []float64) Neuron {
	newCell := new(Cell)
	newCell.weights = genes
	newCell.character = c.character
	newCell.size = c.size
	return *newCell
}

//Breed creates a new cell by mixing the genes of this cell and a pairedCell,
//along with mutations based on the MutationRate
func (c Cell) Breed(pairedCell Breeder) (Neuron, error) {
	if c.size != pairedCell.NumberOfGenes() {
		return *new(Neuron), NewMismatchError(c.size, pairedCell.NumberOfGenes())
	}
	return c.character.NeuronBreeder(c, pairedCell), nil
}

//Clone creates a new Neuron with the exact same behaviour as the original
func (c Cell) Clone() Neuron {
	return Cell{c.weights, c.size, c.character}
}

//NumberOfGenes returns the number of the inputs on a cell
func (c Cell) NumberOfGenes() int {
	return c.size
}

//NthGene returns the nth weight in a cell
func (c Cell) NthGene(n int) float64 {
	return c.weights[n]
}

//NewCell creates a new instance of Cell
func NewCell(size int, character *CellCharacter) *Cell {

	newWeights := make([]float64, size)
	for i := range newWeights {
		newWeights[i] = character.GeneCreator()
	}
	return &Cell{newWeights, size, character}
}
