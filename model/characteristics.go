package model

//GeneCreator is the standard signiture of a method that creates a brand new
//gene
type GeneCreator func() float64

//NeuronBreeder is the standard signature of a method used during the breeding
//process of two networks
type NeuronBreeder func(cellA Neuron, cellB Breeder) Neuron

//Activator determines the output of a unit (normally a value between -1 and 1)
//given an input
type Activator func(float64) float64

//CellCharacter sets the behavior of all cells
type CellCharacter struct {
	Activator     Activator
	GeneCreator   GeneCreator
	NeuronBreeder NeuronBreeder
}
