/*
Author : Dongpin Oh
Date : 12 02 2019
Description : It returns empirical statistic from floating point data that
	represented as golang 1-dimensional slice or 2-dimensional slice
	consisting of (dimensions * values). It includes correlation coefficients,
	covariance, variance, mean.
*/


package	main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
)

type matrix [][]float64

func empCor(x matrix) matrix{
	/*
	* 'x' is an i*j matrix that each row represnts data points(j) and
	* each column represents dimensions (i).
	*
	* It returns an i*i correlation matrix.
	*/
	cov := empCovar(x)
	cor := make([][]float64, len(cov))
	for i := 0; i < len(cor) ; i ++ {
		cor[i] = make([]float64, len(cov[i]))
	}
	
	for i := 0; i < len(cov) ; i ++ {
		for j := i; j < len(cov[i]) ; j++ {
			cor[i][j] = cov[i][j] / (math.Sqrt(cov[i][i]*cov[j][j]))
			cor[j][i] = cor[i][j]
		}
	}
	return cor
}

func empCovar(x matrix) matrix {
	/*
	* 'x' is an i*j matrix that each row represnts data points(j) and
	* each column represents dimensions (i).
	* 
	* It returns an i*i covariance matrix.
	*/
	mean := empMeanMultivar(x)
	cov := make([][]float64 , len(x))
	sum := 0.0

	for i := 0 ; i < len(cov) ; i ++ {
		cov[i] = make([]float64, len(x))
	}
	for i := 0 ; i < len(cov) ; i++ {
		for j := i ; j < len(cov) ; j++ {
			sum = 0.0
			for k := 0 ; k < len(x[i]) ; k++ {
				sum += (x[i][k] - mean[i])*(x[j][k] - mean[j])
			}
			cov[i][j] = sum/float64(len(x[i]))
			cov[j][i] = cov[i][j]
		}

	}
	return cov
}

func empVar(x []float64) float64 {
	mean := empMeanUnivar(x)
	sum := 0.0
	i := 0
	for i < len(x){
		sum += math.Pow((x[i] - mean),2)
		i++
	}
	return sum/float64(len(x))
}

func empMeanMultivar(x matrix) []float64 {
	sum := make([]float64, len(x))
	for i := 0 ; i < len(x); i++{
		for j := 0 ; j < len(x[i]) ; j++{
			sum[i] += x[i][j]
		}
		sum[i] = sum[i]/float64(len(x[i]))
	}
	return sum
}

func empMeanUnivar(x []float64) float64{
	sum := 0.0
	i := 0
	for i < len(x){
		sum += x[i]
		i++
	}
	return sum/float64(len(x))
}

func randomMatrix(n int, m int) [][]float64{
	rand.Seed(time.Now().UTC().UnixNano())
	multiFloat := make([][]float64, n)
	
	for i := 0 ; i < n ; i++ {
		multiFloat[i] = make([]float64, m)
		for j := 0 ; j < m ; j++ {
			multiFloat[i][j] = rand.Float64()*100
		}
	} 
	return multiFloat
}

func main(){
	var randomFloat [100]float64
	multiFloat := randomMatrix(3,100)

	i := 0
	for i < 100 {
		randomFloat[i] = rand.Float64()*100
		i++
	}
	fmt.Println("EmpMean : " , empMeanUnivar(randomFloat[:]))
	fmt.Println("EmpVar : " , empVar(randomFloat[:]))

	fmt.Println("Emp mean Multivar : ", empMeanMultivar(multiFloat))
	fmt.Println("Emp covariance Multivar : ", empCovar(multiFloat))
	fmt.Println("Emp correlation coefficient Multivar : ", empCor(multiFloat))
	fmt.Println("Emp var for 0`th data : ", empVar(multiFloat[0]))
	fmt.Println("Emp var for 1`th data : ", empVar(multiFloat[1]))
	fmt.Println("Emp var for 2`th data : ", empVar(multiFloat[2]))
}