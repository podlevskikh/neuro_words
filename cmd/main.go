package main

import (
	"neuro_word/model/neuro"
	"fmt"
	"math"
	"math/rand"
)

func main() {

	/*data, _ := ioutil.ReadFile("data/train-images-idx3-ubyte")

	fmt.Println(string(data))
	return*/

	fmt.Println("Creating inputs and weights ...")

	net := neuro.NewNet(28*28, []int{16, 16, 3})
	//net.Print()
	r := rand.New(rand.NewSource(99))

	for i := 0; i < 1000000; i++ {
		//fmt.Println("square")
		radius := r.Int31n(5) + 5
		x := r.Int31n(27 - radius)
		y := r.Int31n(27 - radius)
		learnSquare := make([]float64, 0, 28*28)
		for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
				if (i == int(x) || i == int(x+radius)) && j >= int(y) && j <= int(y+radius) ||
					(j == int(y) || j == int(y+radius)) && i >= int(x) && i <= int(x+radius) {
					learnSquare = append(learnSquare, 1)
				} else {
					learnSquare = append(learnSquare, 0)
				}
			}
		}
		net.Learn(learnSquare, []float64{1, 0, 0}, i)
		//fmt.Println(net.Result(learnSquare))

		//fmt.Println("palka_x")
		length_x := r.Int31n(5) + 5
		x1 := r.Int31n(27 - length_x)
		y1 := r.Int31n(27 - length_x)
		learnPalka_x := make([]float64, 0, 28*28)
		for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
				if (i == int(x1)) && j >= int(y1) && j <= int(y1+length_x) {
					learnPalka_x = append(learnPalka_x, 1)
				} else {
					learnPalka_x = append(learnPalka_x, 0)
				}
			}
		}
		net.Learn(learnPalka_x, []float64{0, 1, 0}, i)
		//fmt.Println(net.Result(learnPalka_x))

		//fmt.Println("palka_y")
		length_y := r.Int31n(5) + 5
		x2 := r.Int31n(27 - length_y)
		y2 := r.Int31n(27 - length_y)
		learnPalka_y := make([]float64, 0, 28*28)
		for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
				if (j == int(y2)) && i >= int(x2) && i <= int(x2+length_y) {
					learnPalka_y = append(learnPalka_y, 1)
				} else {
					learnPalka_y = append(learnPalka_y, 0)
				}
			}
		}
		net.Learn(learnPalka_y, []float64{0, 0, 1}, i)
		//fmt.Println(net.Result(learnPalka_y))

		//net.Print()
		/*		net.Learn([]float64{0, 1, 0, 0, 1, 0, 0, 1, 0}, []float64{0, 0, 1})
				net.Learn([]float64{0, 1, 0, 0, 0, 0, 0, 1, 0}, []float64{0, 0, 1})
				//net.Print()
				net.Learn([]float64{0, 0, 1, 0, 0, 0, 0, 0, 1}, []float64{1, 0, 0})
				net.Learn([]float64{0, 0, 1, 0, 0, 1, 0, 0, 1}, []float64{1, 0, 0})*/
		//net.Print()
	}

	hashtag := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 10 || i == 20 || j == 10 || j == 20 {
				hashtag = append(hashtag, 1)
			} else {
				hashtag = append(hashtag, 0)
			}
		}
	}
	fmt.Println("hashtag")
	fmt.Println(net.Result(hashtag))

	fatcross := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 13 || i == 24 || j == 13 || j == 24 {
				fatcross = append(fatcross, 1)
			} else {
				fatcross = append(fatcross, 0)
			}
		}
	}
	fmt.Println("fatcross")
	fmt.Println(net.Result(fatcross))

	cross := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 13 || j == 13 {
				cross = append(cross, 1)
			} else {
				cross = append(cross, 0)
			}
		}
	}
	fmt.Println("cross")
	fmt.Println(net.Result(cross))

	fatcorner := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 0 || i == 1 || j == 0 || j == 1 {
				fatcorner = append(fatcorner, 1)
			} else {
				fatcorner = append(fatcorner, 0)
			}
		}
	}
	fmt.Println("fatcorner")
	fmt.Println(net.Result(fatcorner))

	corner := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 0 || j == 0 {
				corner = append(corner, 1)
			} else {
				corner = append(corner, 0)
			}
		}
	}
	fmt.Println("corner")
	fmt.Println(net.Result(corner))

	longline := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 0 {
				longline = append(longline, 1)
			} else {
				longline = append(longline, 0)
			}
		}
	}
	fmt.Println("longline")
	fmt.Println(net.Result(longline))

	fatline := make([]float64, 0, 28*28)
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 20 && j < 10 && j > 5 {
				fatline = append(fatline, 1)
			} else {
				fatline = append(fatline, 0)
			}
		}
	}
	fmt.Println("fatline")
	fmt.Println(net.Result(fatline))

	/*firstLayer := []neuro.Neuron{{
		Name:         "layer 1 cell 1",
		Weight:       []float64{0.2, 0.2, 0.2},
		OffsetWeight: 0.8,
	}, {
		Name:         "layer 1 cell 2",
		Weight:       []float64{0.4, 0.4, 0.4},
		OffsetWeight: 0.8,
	}, {
		Name:         "layer 1 cell 3",
		Weight:       []float64{0.6, 0.6, 0.6},
		OffsetWeight: 0.8,
	}}
	secondLayer := []neuro.Neuron{{
		Name:         "layer 2 cell 1",
		Weight:       []float64{0.5, 0.5, 0.5},
		OffsetWeight: 0.2,
	}}

	input := []float64{1.5, 2.0, 3.0}

	firstLayerOutput := make([]float64, 0, len(firstLayer))
	for _, firstLayerNeuron := range firstLayer {
		firstLayerOutput = append(firstLayerOutput, firstLayerNeuron.Evaluate(input))
	}
	fmt.Print(firstLayerOutput)
	secondLayerOutput := make([]float64, 0, len(secondLayer))
	for _, secondLayerNeuron := range secondLayer {
		secondLayerOutput = append(secondLayerOutput, secondLayerNeuron.Evaluate(firstLayerOutput))
	}

	fmt.Printf("result %f\n", secondLayerOutput[0])*/
}

func train(trials int, inputs []float64, weights []float64, desired float64, learningRate float64) {

	for i := 1; i < trials; i++ {
		weights = learn(inputs, weights, learningRate)
		output := evaluate(inputs, weights)
		errorResult := evaluateError(desired, output)

		fmt.Print("Output: ")
		fmt.Print(math.Round(output*100) / 100)
		fmt.Print("\nError: ")
		fmt.Print(math.Round(errorResult*100) / 100)
		fmt.Print("\n\n")
	}

}

func learn(inputVector []float64, weightVector []float64, learningRate float64) []float64 {
	for index, inputValue := range inputVector {
		if inputValue > 0.00 {
			weightVector[index] = weightVector[index] + learningRate
		}
	}

	return weightVector
}

func evaluate(inputVector []float64, weightVector []float64) float64 {
	result := 0.00

	for index, inputValue := range inputVector {
		layerValue := inputValue * weightVector[index]
		result = result + layerValue
	}

	return result
}

func evaluateError(desired float64, actual float64) float64 {
	return desired - actual
}
