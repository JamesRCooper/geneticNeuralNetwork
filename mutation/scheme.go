package mutation

import (
	"math"
	"math/rand"

	m "github.com/JamesRCooper/geneticNeuralNetwork/model"
)

//BuildGaussianBreeder creates a gaussian breed that uses the standard method of
//choosing between two genes or creating a random third gene. The chances of a
//third gene are based upon the supplied mutation rate. If a gene is chosen, a
//Gaussian distributed value is added based upon the supplied standard deviation.
func BuildGaussianBreeder(
	mutationRate float64, stdDev float64, creator m.GeneCreator) m.NeuronBreeder {

	return func(cellA m.Neuron, cellB m.Breeder) m.Neuron {
		genes := make([]float64, cellA.NumberOfGenes())
		for index := 0; index < cellA.NumberOfGenes(); index++ {
			cellChoice := rand.Float64()
			if cellChoice < mutationRate {
				genes[index] = creator()
			} else if cellChoice < (mutationRate/2.0)+0.5 {
				geneAugment := rand.NormFloat64() * stdDev
				genes[index] = cellA.NthGene(index) + geneAugment
			} else {
				geneAugment := rand.NormFloat64() * stdDev
				genes[index] = cellB.NthGene(index) + geneAugment
			}
		}

		return cellA.CreateNew(genes)
	}
}

//BuildNormalBreeder creates a Normal breeder that uses the standarsd method of
//choosing between two genes, or creating a new random third gene. The chances
//of a new gene are based upon the mutation rate; the generation of the new gene
//is perform with the GeneCreator
func BuildNormalBreeder(
	mutationRate float64, creator m.GeneCreator) m.NeuronBreeder {

	return func(cellA m.Neuron, cellB m.Breeder) m.Neuron {
		genes := make([]float64, cellA.NumberOfGenes())
		for index := 0; index < cellA.NumberOfGenes(); index++ {
			cellChoice := rand.Float64()
			if cellChoice < mutationRate {
				genes[index] = creator()
			} else if cellChoice < (mutationRate/2.0)+0.5 {
				genes[index] = cellA.NthGene(index)
			} else {
				genes[index] = cellB.NthGene(index)
			}
		}

		return cellA.CreateNew(genes)
	}
}

//LogisticSigmoidActivator uses an augmetned version of the logistics function
//to return a value between -1 and 1, with prejudice towards the boundaries
func LogisticSigmoidActivator(input float64) float64 {
	return (2.0 / (1.0 + math.Exp(-1.0*input))) - 1.0
}

//StandardGeneCreator creates a new gene with a random value noramlly
//distributed between -1 and 1
func StandardGeneCreator() float64 {
	return (2.0 * rand.Float64()) - 1.0
}

//NormalyDistributedGeneCreator cretes a new gene with a normal distribution:
//average=0 and standard deviation of 1
func NormalyDistributedGeneCreator() float64 {
	return rand.NormFloat64()
}
