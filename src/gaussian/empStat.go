/*
Author : Dongpin Oh
Date : 12 02 2019
Description : It returns empirical statistic from floating point data that
	represented as golang 1-dimensional slice or 2-dimensional slice
	consisting of (dimensions * values). It includes correlation coefficients,
	covariance, variance, mean.
*/

package gaussian 

import (
	"math/rand"
	"math"
	"time"
	"matop"
)

type Matrix = matop.Matrix 
type Vector = matop.Vector

func EmpCor(x Matrix) Matrix{
	/*
	* 'x' is an i*j Matrix that each row represnts data points(i) and
	* each column represents dimensions (j).
	*
	* It returns an i*i correlation Matrix.
	*/
	cov := EmpCovar(x)
	cor := make(Matrix, len(cov))
	for i := 0; i < len(cor) ; i ++ {
		cor[i] = make(Vector, len(cov[i]))
	}
	
	for i := 0; i < len(cov) ; i ++ {
		for j := i; j < len(cov[i]) ; j++ {
			cor[i][j] = cov[i][j] / (math.Sqrt(cov[i][i]*cov[j][j]))
			cor[j][i] = cor[i][j]
		}
	}
	return cor
}

func EmpCovar(x Matrix) Matrix {
	/*
	* 'x' is an i*j Matrix that each column represnts data points(j) and
	* each row represents dimensions (i). (Note that each column Vector
	* represents data)
	* 
	* It returns an i*i covariance Matrix.
	*/
	mean := EmpMeanMultivar(x)
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

func EmpVar(x Vector) float64 {
	mean := EmpMeanUnivar(x)
	sum := 0.0
	i := 0
	for i < len(x){
		sum += math.Pow((x[i] - mean),2)
		i++
	}
	return sum/float64(len(x))
}

func EmpMeanMultivar(x Matrix) []float64 {
	sum := make([]float64, len(x))
	for i := 0 ; i < len(x); i++{
		for j := 0 ; j < len(x[i]) ; j++{
			sum[i] += x[i][j]
		}
		sum[i] = sum[i]/float64(len(x[i]))
	}
	return sum
}

func EmpMeanUnivar(x Vector) float64{
	sum := 0.0
	i := 0
	for i < len(x){
		sum += x[i]
		i++
	}
	return sum/float64(len(x))
}

func RandomMatrix(n int, m int) Matrix{
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