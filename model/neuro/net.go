package neuro

import (
	"fmt"
	"math/rand"
)

type Net struct {
	Layers [][]*Neuron
}

func NewNet(inputCount int, layersCount []int) *Net {
	/*w1 :=
	return &Net{Layers: [][]*Neuron{
		{{}},
	}}*/

	r := rand.New(rand.NewSource(99))

	n := Net{}
	n.Layers = make([][]*Neuron, 0, len(layersCount))
	for _, layerCount := range layersCount {
		n.Layers = append(n.Layers, make([]*Neuron, layerCount))
	}
	for i := range n.Layers {
		for j := range n.Layers[i] {
			nCurrent := &Neuron{I: i, J: j}
			if i+1 < len(n.Layers) {
				nCurrent.Outputs = make([]*Weight, 0, len(n.Layers[i+1]))
			}
			if i == 0 {
				nCurrent.Inputs = make([]*Weight, 0, inputCount)
				for k := 0; k < inputCount; k++ {
					w := Weight{
						Value: r.Float64(),
					}
					nCurrent.Inputs = append(nCurrent.Inputs, &w)
				}
			} else {
				nCurrent.Inputs = make([]*Weight, 0, len(n.Layers[i-1]))
				for _, nPrev := range n.Layers[i-1] {
					w := Weight{
						Value: r.Float64(),
					}
					nPrev.Outputs = append(nPrev.Outputs, &w)
					nCurrent.Inputs = append(nCurrent.Inputs, &w)
				}
			}
			n.Layers[i][j] = nCurrent
		}
	}
	return &n
}

func (n *Net) Learn(input []float64, checkResult []float64, i int) {
	nextLayerInput := input
	for _, layer := range n.Layers {
		layerResult := make([]float64, 0, len(layer))
		for _, neuron := range layer {
			layerResult = append(layerResult, neuron.Evaluate(nextLayerInput))
		}
		nextLayerInput = layerResult
	}

	err := 0.0
	for i, res := range checkResult {
		err += (res - nextLayerInput[i]) * (res - nextLayerInput[i])
	}
	if err / float64(len(checkResult)) > 0.001 {
		fmt.Printf("%d %f\n", i, err / float64(len(checkResult)))
	}

	var sigmas []float64
	for i := len(n.Layers) - 1; i >= 0; i-- {
		if i == len(n.Layers) - 1 {
			for j, neuron := range n.Layers[i] {
				sigmas = append(sigmas, neuron.CalculateSigmaLastLayer(checkResult[j]))
			}
		} else {
			currentSigmas := make([]float64, 0, len(n.Layers[i]))
			for _, neuron := range n.Layers[i] {
				currentSigmas = append(sigmas, neuron.CalculateSigma(sigmas))
			}
			sigmas = currentSigmas
		}
	}

	for _, layer := range n.Layers {
		for _, neuron := range layer {
			neuron.Spread()
		}
	}
}

func (n *Net) Result(input []float64) []float64 {
	nextLayerInput := input
	for _, layer := range n.Layers {
		layerResult := make([]float64, 0, len(layer))
		for _, neuron := range layer {
			layerResult = append(layerResult, neuron.Evaluate(nextLayerInput))
		}
		nextLayerInput = layerResult
	}
	return nextLayerInput
}

func (n *Net) Print() {
	fmt.Println("<table border=\"1\">")
	fmt.Println("<tr>")
	for _, layer := range n.Layers {
		fmt.Println("<td>")
		fmt.Println("<table>")
		for _, neuron := range layer {
			fmt.Println("<tr>")
			fmt.Println("<td>")
			neuron.Print()
			fmt.Println("</td>")
			fmt.Println("</tr>")
		}
		fmt.Println("</table>")
		fmt.Println("</td>")
	}
	fmt.Println("</tr>")
	fmt.Println("</table>")
}
