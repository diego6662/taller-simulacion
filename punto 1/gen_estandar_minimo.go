package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var a float64
	var m float64
	var xn float64
	fmt.Print("ingrese a:")
	fmt.Scanln(&a)
	fmt.Print("ingrese m:")
	fmt.Scanln(&m)
	fmt.Print("ingrese x0:")
	fmt.Scanln(&xn)
	var Rn = []float64{}
	var Xn = []int{}
	for {
		noRepeat := true
		xn = math.Mod((a * xn), m)
		Xn = append(Xn, int(xn))
		rn := xn / m
		size := len(Rn)
		for i := 0; i < size; i++ {
			if rn == Rn[i] {
				noRepeat = false
				break
			}
		}
		if !noRepeat {
			break
		}
		Rn = append(Rn, rn)
	}
	file, err := os.Create("Tabla.txt")
	file.WriteString("---------------TABLA---------------\n")
	file.WriteString("N || Xn || Rn\n")
	file.WriteString("-----------------------------------\n")
	for i := 0; i < len(Rn); i++ {
		str := fmt.Sprintf("%d || %d || %f\n", i, Xn[i], Rn[i])
		file.WriteString(str)
		file.WriteString("-----------------------------------\n")
		if i == len(Rn)-1 {
			fmt.Printf("El periodo es de longitud %d \n", i+1)
		}
	}
	file.Close()
	number, err := os.Create("Number.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(Rn); i++ {
		str := fmt.Sprintf("%f", Rn[i])
		number.WriteString(str)
		number.WriteString(",")
	}
	number.Close()

}
