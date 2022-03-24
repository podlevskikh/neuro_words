package neuro

import (
	"fmt"
	"math"
)

const Step = 0.2

type Neuron struct {
	I            int
	J            int
	Inputs       []*Weight
	Outputs      []*Weight
	InputComes   []float64
	OutputResult float64
	Delta        float64
	Offset       float64
}

type Weight struct {
	Value float64
}

func (n *Neuron) Evaluate(input []float64) float64 {
	n.InputComes = input

	result := n.Offset

	for index, inputValue := range input {
		result += inputValue * n.Inputs[index].Value
	}

	n.OutputResult = 1 / (1 + math.Exp(-result))
	return n.OutputResult
}

func (n *Neuron) CalculateSigmaLastLayer(checkOutput float64) float64 {
	n.Delta = (0 - n.OutputResult) * (1 - n.OutputResult) * (checkOutput - n.OutputResult)
	return n.Delta
}

func (n *Neuron) CalculateSigma(sigmas []float64) float64 {
	sigmaWightSum := 0.0
	for index, sigma := range sigmas {
		sigmaWightSum += sigma * n.Outputs[index].Value
	}
	n.Delta = n.OutputResult * (1 - n.OutputResult) * sigmaWightSum
	return n.Delta
}

func (n *Neuron) Spread() {
	for i, weight := range n.Inputs {
		weight.Value -= Step * n.Delta * n.InputComes[i]
	}
	n.Offset -= Step * n.Delta
}

func (n *Neuron) Print() {
	fmt.Println("<table>")

	fmt.Println("<tr>")
	fmt.Println("<td>")
	for _, in := range n.Inputs {
		fmt.Println(in.Value)
		fmt.Println("</br>")
	}
	fmt.Println("</td>")

	fmt.Println("<td>")
	fmt.Printf("%d, %d, %f, %f", n.I, n.J, n.OutputResult, n.Delta)
	fmt.Println("</td>")

	fmt.Println("<td>")
	for _, o := range n.Outputs {
		fmt.Println(o.Value)
		fmt.Println("</br>")
	}
	fmt.Println("</td>")
	fmt.Println("</tr>")

	fmt.Println("</table>")
}
