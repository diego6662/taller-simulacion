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
	fe := float64(len(nums)) / 10.0 //  frecuencia esperada

	chiCrit := 16.92
	var fo = [10]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range nums {
		switch {
		case v >= 0.0 && v < 0.1:
			fo[0]++
		case v >= 0.1 && v < 0.2:
			fo[1]++
		case v >= 0.2 && v < 0.3:
			fo[2]++
		case v >= 0.3 && v < 0.4:
			fo[3]++
		case v >= 0.4 && v < 0.5:
			fo[4]++
		case v >= 0.5 && v < 0.6:
			fo[5]++
		case v >= 0.6 && v < 0.7:
			fo[6]++
		case v >= 0.7 && v < 0.8:
			fo[7]++
		case v >= 0.8 && v < 0.9:
			fo[8]++
		case v >= 0.9 && v <= 1:
			fo[9]++
		}
	}
	chiCal := 0.0
	for i := 0; i < len(fo); i++ {
		chiCal += (math.Pow((fe - fo[i]), 2)) / fe
	}
	fmt.Printf("chi calculado:%f\n", chiCal)
	if chiCal <= chiCrit {
		fmt.Println("cumple la prueba de uniformidad")
	} else {
		fmt.Println("no cumple la prueba de uniformidad")
	}
	tabla, err := os.Create("Tabla.txt")
	if err != nil {
		log.Fatal(err)
	}
	tabla.WriteString("--------------------tabla-------------------------\n")
	tabla.WriteString("--------------------------------------------------\n")
	tabla.WriteString("    rango     || FO ||    FE   || (FE - FO)²/FE ||\n")
	intInf := 0.0
	intSup := 1.0 / 10.0
	incremento := 1.0 / 10.0
	for i := 0; i < len(fo); i++ {
		str := fmt.Sprintf("%0.3f - %0.3f || %0.0f || %0.4f ||        %0.4f ||\n", intInf, intSup, fo[i], fe, math.Pow((fe-fo[i]), 2)/fe)
		tabla.WriteString(str)
		tabla.WriteString("--------------------------------------------------\n")
		intInf = intSup
		intSup += incremento
	}
	tabla.Close()
}
