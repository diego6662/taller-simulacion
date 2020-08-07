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
	numCorr := 0.0 //numero de corrdas
	np := 0.0      //numero de crecimientos
	nd := 0.0      //numero de decrementos
	//nivelConf := 0.05 //nivel de confianza
	value := nums[0]
	var up bool
	tabla, err := os.Create("tabla.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {

			tabla.WriteString("*")

			value = nums[i+1]
			if nums[i] > value {

				up = true

			} else {
				up = false

			}
			value = nums[i]
			continue
		}

		if nums[i] > value {

			if !up {
				np++
				up = true
				numCorr++
				tabla.WriteString(" || ")
			}
			tabla.WriteString(" + ")

		} else {

			if up {
				nd++
				up = false
				numCorr++
				tabla.WriteString(" || ")
			}
			tabla.WriteString(" - ")
		}
		value = nums[i]
	}
	media := (2.0 * float64(len(nums))) / 3.0
	varianza := ((16.0 * float64(len(nums))) - 29) / 90
	zObs := (numCorr - media) / math.Sqrt(varianza)
	intSup := 1.96
	intInf := -1.96

	if zObs >= intInf && zObs <= intSup {
		fmt.Println("cumple la prueba de independencia")
	} else {
		fmt.Println("no cumple la prueba de independencia")
	}
	str := fmt.Sprintf("\nzObs = %f\nnumero de corridas:%0.0f\nnumero subidas:%0.0f\nnumero bajadas:%0.0f", zObs, numCorr, np, nd)
	tabla.WriteString(str)
	tabla.Close()
}
