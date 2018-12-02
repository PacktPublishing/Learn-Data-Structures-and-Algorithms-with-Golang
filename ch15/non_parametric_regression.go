//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing math,sort,os,csv,io,strconv and fmt packages

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

//Distance method between point1 and point2
func Distance(point1 []float64, point2 []float64) float64 {
	var value float64
	value = 0.0
	var i int
	for i, _ = range point1 {
		value = value + math.Pow(point1[i]-point2[i], 2)
	}
	return math.Sqrt(value)
}

//DataSlice
type DataSlice struct {
	sort.Interface
	indices []int
}

// Swap method
func (slice DataSlice) Swap(i int, j int) {
	slice.Interface.Swap(i, j)
	slice.indices[i], slice.indices[j] = slice.indices[j], slice.indices[i]
}

// NewDataSlice method
func NewDataSlice(interf sort.Interface) *DataSlice {
	var slice *DataSlice
	slice = &DataSlice{Interface: interf, indices: make([]int, interf.Len())}
	var i int
	for i = range slice.indices {
		slice.indices[i] = i
	}
	return slice
}

//NewFloat64DataSlice method
func NewFloat64DataSlice(values []float64) *DataSlice { return NewDataSlice(sort.Float64Slice(values)) }

// DataRow class
type DataRow struct {
	colName string
	value   int
}

// DataList class
type DataList []DataRow

// Len method
func (list DataList) Len() int {
	return len(list)
}

// Swap method
func (list DataList) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}

// Less method
func (list DataList) Less(i int, j int) bool {
	if list[i].value == list[j].value {
		return list[i].colName < list[j].colName
	} else {
		return list[i].value > list[j].value
	}
}

// GetDataMap method
func GetDataMap(data []string) map[string]int {
	var dataMap map[string]int
	dataMap = map[string]int{}
	var element string
	for _, element = range data {
		dataMap[element] = dataMap[element] + 1
	}
	return dataMap
}

// Algo class
type Algo struct {
	kValue int
	data   [][]float64
	names  []string
}

// GetCloseness method
func (algo *Algo) GetCloseness(X [][]float64, Y []string) {
	algo.data = X
	algo.names = Y
}

// GetPredictions method
func (algo *Algo) GetPredictions(X [][]float64) []string {

	var predictedValue []string
	predictedValue = []string{}
	var source []float64
	for _, source = range X {
		var (
			list      []float64
			nearNames []string
		)
		var dest []float64
		for _, dest = range algo.data {
			list = append(list, Distance(source, dest))
		}
		var slice *DataSlice
		slice = NewFloat64DataSlice(list)
		sort.Sort(slice)
		var index []int
		index = slice.indices[:algo.kValue]

		var i int
		for _, i = range index {
			nearNames = append(nearNames, algo.names[i])
		}

		var freq map[string]int
		freq = GetDataMap(nearNames)

		var rowList DataList
		rowList = DataList{}
		var str string
		for str, i = range freq {
			var row DataRow
			row = DataRow{str, i}
			rowList = append(rowList, row)
		}
		sort.Sort(rowList)
		predictedValue = append(predictedValue, rowList[0].colName)
	}
	return predictedValue

}

// main method
func main() {

	var matrixData [][]string
	matrixData = [][]string{}

	var file *os.File
	var err error
	file, err = os.Open("loan_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var reader *csv.Reader
	reader = csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		var data []string
		var err error
		data, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		matrixData = append(matrixData, data)
	}

	var X [][]float64
	var Y []string
	X = [][]float64{}
	Y = []string{}

	var data []string
	for _, data = range matrixData {

		var temp []float64
		temp = []float64{}

		var i string
		for _, i = range data[:4] {
			var value float64
			var err error
			value, err = strconv.ParseFloat(i, 64)
			if err != nil {
				panic(err)
			}
			temp = append(temp, value)
		}
		X = append(X, temp)

		Y = append(Y, data[4])

	}

	var (
		trainX [][]float64
		trainY []string
		testX  [][]float64
		testY  []string
	)
	var i int
	for i, _ = range X {
		if i%2 == 0 {
			trainX = append(trainX, X[i])
			trainY = append(trainY, Y[i])
		} else {
			testX = append(testX, X[i])
			testY = append(testY, Y[i])
		}
	}

	var algo Algo
	algo = Algo{}
	algo.kValue = 8
	algo.GetCloseness(trainX, trainY)
	var predicted []string
	predicted = algo.GetPredictions(testX)

	var correct int
	correct = 0
	//var i int
	for i, _ = range predicted {
		if predicted[i] == testY[i] {
			correct += 1
		}
	}
	fmt.Println("correct prediction", correct)
	fmt.Println("predicted number", len(predicted))
	fmt.Println(" corrected/predicted", float64(correct)/float64(len(predicted)))

}
