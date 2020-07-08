package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("../../punto 1/Number.txt")
	if err != nil {
		log.Fatal(err)
	}
	numsStr := strings.Split(string(file), ",")
	var nums = []float64{}
	for i := 0; i < len(numsStr)-1; i++ {
		floatTemp, err := strconv.ParseFloat(numsStr[i], 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, floatTemp)
	}

	//estos son los datos de frecuencia obtenida
	td := 0
	onePair := 0
	terna := 0
	for _, i := range nums {

		var hash map[string]int
		hash = make(map[string]int)
		temp := fmt.Sprintf("%0.3f", i)
		temp = temp[2:]
		for j := 0; j < 3; j++ {
			key := string(temp[j])
			hash[key]++
		}
		size := len(hash)

		switch size {
		case 1:
			terna++
		case 2:
			onePair++
		case 3:
			td++
		}

	}
	size := len(nums)
	tdProb := 0.72 * float64(size)
	onePairPro := 0.27 * float64(size)
	ternaPro := 0.01 * float64(size)
	var FO = [3]int{td, onePair, terna}
	var FE = [3]float64{tdProb, onePairPro, ternaPro}
	var RN = []float64{}
	var intV = [7]string{"todos dif", "1 pareja", "terna"}
	chiCalc := 0.0
	for i := 0; i < len(FO); i++ {
		chiCalc += math.Pow((FE[i]-float64(FO[i])), 2) / FE[i]
		RN = append(RN, chiCalc)
	}
	chiCrit := 6.0
	if chiCalc <= chiCrit {
		fmt.Println("cumple la prueba de independencia")
	} else {
		fmt.Println("no cumple la prueba de independencia")
	}
	tabla, err := os.Create("Tabla.txt")
	tabla.WriteString("--------------------tabla-------------------------\n")

	tabla.WriteString("    rango     || FO ||    FE   || (FE - FO)²/FE ||\n")
	tabla.WriteString("--------------------------------------------------\n")
	for i := 0; i < len(FO); i++ {
		str := fmt.Sprintf("%s  ||  %d||  %0.0f ||        %0.5f ||\n", intV[i], FO[i], FE[i], RN[i])
		tabla.WriteString(str)
		tabla.WriteString("--------------------------------------------------\n")

	}
	tabla.Close()
}
