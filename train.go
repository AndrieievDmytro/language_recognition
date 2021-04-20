package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	testFolder  = "test"
	trainFolder = "train"
	perceptrons = CreatePerceptrons(trainFolder)
	threshold   = 0.01
)

func CreatePerceptrons(folder string) []Perceptron {
	var arr []Perceptron
	for i := range AssignPercName(folder) {
		// fmt.Println(GenerateWeights())
		var per = Perceptron{AssignPercName(folder)[i], threshold, GenerateWeights()}
		arr = append(arr, per)
	}
	return arr
}

func Train(desiredAcc float64) {
	var accurancy float64 = 0.0
	m := FileLetterPropirtion(trainFolder)
	for accurancy < desiredAcc {
		for folderName, folders := range m {
			for _, letterProp := range folders {
				for idx, perceptron := range perceptrons {
					output := perceptron.Predict(letterProp)
					var desirVal float64
					if perceptron.Label != folderName {
						desirVal = 0.0
					} else {
						desirVal = 1.0
					}
					perceptrons[idx].DeltaRule(letterProp, desirVal, output)
				}
			}
		}
		accurancy = CalculateAccuracy()
		fmt.Println(accurancy)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text: ")
	langText, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println("The text is in", PredictTest(GetLetterProportion(langText)))
}

func PredictTest(arr []float64) string {
	results := make(map[string]float64)
	for _, perceptron := range perceptrons {
		results[perceptron.Label] = perceptron.Predict(arr)
	}
	names := make([]string, 0, len(results))
	for name := range results {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return results[names[i]] > results[names[j]]
	})
	first := names[0]
	return first
}

func CalculateAccuracy() float64 {
	counter := 0.0
	fileCount := 0.0
	m := FileLetterPropirtion(testFolder)
	for folderName, files := range m {
		for _, letterProp := range files {
			fileCount++
			predictionResult := PredictTest(letterProp)
			if predictionResult == folderName {
				counter++
			}
		}
	}
	return (counter / fileCount) * 100.0
}
