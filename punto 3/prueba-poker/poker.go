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
	twoPair := 0
	terna := 0
	ternaPair := 0
	poker := 0
	quintilla := 0
	for _, i := range nums {

		var hash map[string]int
		hash = make(map[string]int)
		temp := fmt.Sprintf("%0.5f", i)
		temp = temp[2:]

		for j := 0; j < 5; j++ {
			key := string(temp[j]) //12341
			hash[key]++
		}
		size := len(hash)

		switch size {
		case 1:
			quintilla++
		case 2:
			max := 0
			for _, value := range hash {
				if value > max {
					max = value
				}

			}
			if max == 4 {
				poker++

			} else {
				ternaPair++

			}
		case 3:
			max := 0
			for _, value := range hash {
				if value > max {
					max = value
				}
			}
			if max == 3 {
				terna++ //11134

			} else {
				twoPair++ //11225

			}
		case 4:
			onePair++
		case 5:
			td++

		}

	}
	size := len(nums)
	tdProb := 0.3024 * float64(size)
	onePairPro := 0.5040 * float64(size)
	twoPairPro := 0.1080 * float64(size)
	ternaPairPro := 0.0090 * float64(size)
	ternaPro := 0.0720 * float64(size)
	pokerPro := 0.0045 * float64(size)
	quintillaPro := 0.0001 * float64(size)
	var FO = [7]int{td, onePair, twoPair, ternaPair, terna, poker, quintilla}
	var FE = [7]float64{tdProb, onePairPro, twoPairPro, ternaPairPro, ternaPro, pokerPro, quintillaPro}
	var RN = []float64{}
	var intV = [7]string{"todos dif", "1 pareja", "2 parejas", "terna par", "terna", "poker", "quintilla"}
	chiCalc := 0.0
	for i := 0; i < len(FO); i++ {
		chiCalc += math.Pow((FE[i]-float64(FO[i])), 2) / FE[i]
		RN = append(RN, chiCalc)
	}
	chiCrit := 12.5916
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
