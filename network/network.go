package network


type Network struct {
	InputNodes int
	HiddenNodes int
	OutputNodes int

	LearningRate float32

}

func (n *Network) train() {
	return
}

func (n *Network) query() {
	return
}

func New(inputNodes, hiddenNodes, outputNodes int, learningRate float32) *Network {
	n := &Network{InputNodes:inputNodes, HiddenNodes:hiddenNodes, OutputNodes:outputNodes, LearningRate:learningRate}
	return n
}
