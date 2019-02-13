package main

import (
	"fmt"
	"math/rand"
	"gaussian"
)

func main(){
	var randomFloat [100]float64
	multiFloat := gaussian.RandomMatrix(3,100)

	i := 0
	for i < 100 {
		randomFloat[i] = rand.Float64()*100
		i++
	}
	fmt.Println("EmpMean : " , gaussian.EmpMeanUnivar(randomFloat[:]))
	fmt.Println("EmpVar : " , gaussian.EmpVar(randomFloat[:]))

	fmt.Println("Emp mean Multivar : ", gaussian.EmpMeanMultivar(multiFloat))
	fmt.Println("Emp covariance Multivar : ", gaussian.EmpCovar(multiFloat))
	fmt.Println("Emp correlation coefficient Multivar : ", gaussian.EmpCor(multiFloat))
	fmt.Println("Emp var for 0`th data : ", gaussian.EmpVar(multiFloat[0]))
	fmt.Println("Emp var for 1`th data : ", gaussian.EmpVar(multiFloat[1]))
	fmt.Println("Emp var for 2`th data : ", gaussian.EmpVar(multiFloat[2]))
}