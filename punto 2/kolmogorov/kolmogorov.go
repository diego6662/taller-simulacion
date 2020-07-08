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
	n := len(nums) //numero de elementos

	c := 10.0 //numero de clases
	gl := n   //grado de libertad
	inc := 1.0 / c

	var FO = [10]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range nums {
		switch {
		case v >= 0.0 && v < 0.1:
			FO[0]++
		case v >= 0.1 && v < 0.2:
			FO[1]++
		case v >= 0.2 && v < 0.3:
			FO[2]++
		case v >= 0.3 && v < 0.4:
			FO[3]++
		case v >= 0.4 && v < 0.5:
			FO[4]++
		case v >= 0.5 && v < 0.6:
			FO[5]++
		case v >= 0.6 && v < 0.7:
			FO[6]++
		case v >= 0.7 && v < 0.8:
			FO[7]++
		case v >= 0.8 && v < 0.9:
			FO[8]++
		case v >= 0.9 && v <= 1:
			FO[9]++
		}
	}
	var FOA = []float64{}
	foa := 0.0
	for i := 0; i < len(FO); i++ {
		foa += FO[i]
		FOA = append(FOA, foa)
	}
	var POA = []float64{}
	for i := 0; i < len(FOA); i++ {
		temp := FOA[i] / float64(n)
		POA = append(POA, temp)
	}
	var PEA = []float64{}
	pro := 1.0 / c

	for i := 0; i < len(FO); i++ {
		PEA = append(PEA, pro)
		pro += inc
	}
	//calculamos el dmcrit
	DMcrit := 0.0
	if gl < 35 {
		fmt.Printf("Ingrese el valor de DMcrit tabulado en la tabla de kolmogorov para %d grados de libertad y un nivel de confianza de 0.05:", gl)
		fmt.Scan(&DMcrit)
	} else {
		DMcrit = 1.36 / math.Sqrt(float64(n))
	}
	DMcalc := 0.0
	for i := 0; i < len(PEA); i++ {
		temp := math.Abs(PEA[i] - POA[i])
		if temp > DMcalc {
			DMcalc = temp
		}
	}
	fmt.Println(DMcalc)
	if DMcalc <= DMcrit {
		fmt.Println("cumple la prueba de uniformidad")
	} else {
		fmt.Println("no cumple la prueba de uniformidad")
	}
	intInf := 0.0
	intSup := 1.0 / c
	tabla, err := os.Create("Tabla.txt")
	tabla.WriteString("----------------------TABLA------------------------------------\n")
	tabla.WriteString("     RANGO    || FO || FOA || POA   ||    PEA || |PEA - POA| ||\n")
	tabla.WriteString("----------------------------------------------------------------\n")

	for i := 0; i < len(FO); i++ {
		if DMcalc == math.Abs(PEA[i]-POA[i]) {
			tabla.WriteString("*")
		}
		str := fmt.Sprintf("%0.3f - %0.3f || %0.0f || %0.0f || %0.4f || %0.4f || %f ||\n", intInf, intSup, FO[i], FOA[i], POA[i], PEA[i], math.Abs(PEA[i]-POA[i]))
		tabla.WriteString(str)
		tabla.WriteString("----------------------------------------------------------------\n")
		intInf = intSup
		intSup += inc
	}
	tabla.Close()
}
