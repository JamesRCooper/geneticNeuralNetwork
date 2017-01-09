package model

//CellCharacter sets the behavior of all cells
type CellCharacter struct {
	MutationRate float64
	Activater    func(float64) float64
	GeneCreator  func() float64
}
