package model

//Stack is general term that can apply to layer, a subdomain, or a network
type Stack interface {
	CalculateOutput(inputs []float64) ([]float64, error)
}

//Network includes inputs, layers, and the output
type Network struct {
	layers      []Layer
	numOfLayers int
}

//CalculateOutput calculates the ouputs of all sub domains
func (n *Network) CalculateOutput(inputs []float64) ([]float64, error) {
	inputsToNextLayer := inputs
	var err error
	for _, layer := range n.layers {
		inputsToNextLayer, err = layer.CalculateOutput(inputsToNextLayer)
		if err != nil {
			return inputsToNextLayer, err
		}
	}
	return inputsToNextLayer, nil
}

//NewNetwork generates a network or sub network for use in a larger network
func NewNetwork(
	numberOfInputs int, sizeOfLayers []int, cellCharacter *CellCharacter) Network {

	numberOfLayers := len(sizeOfLayers)
	networkLayers := make([]Layer, 0, numberOfLayers)
	previousLayerSize := numberOfInputs
	for _, layerSize := range sizeOfLayers {
		networkLayers = append(
			networkLayers, NewLayer(layerSize, previousLayerSize, cellCharacter))
		previousLayerSize = layerSize
	}
	return Network{networkLayers, numberOfLayers}
}

//Breed the host network with a partner network using an equal split
//of genes, as well as using the supplied mutation rate
func (n *Network) Breed(partnerNetwork *Network) {
	for index := 0; index < len(n.layers); index++ {
		n.layers[index].Breed(partnerNetwork.layers[index])
	}
}
