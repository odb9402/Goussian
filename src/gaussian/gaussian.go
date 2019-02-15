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
	"fmt"
	"math"
	"matop"
	"errors"
)

const pi = math.Pi

func GaussianUniv(x float64, mean float64, variance float64) float64{
	return 1/math.Sqrt(2*pi*math.Abs(variance))*math.Exp((x - mean)/(2*math.Abs(variance)))
}

func GaussianMultiv(x Vector, mean Vector, cov Matrix) float64{
	if len(x) != len(mean){
		message := fmt.Sprintf("ERROR: Length of input vector (%d) and mean (%d) are different.\n", len(x),len(mean))
		errors.New(message)
		fmt.Println(errors.New(message))
		return 0.0
	}
	a := math.Sqrt(matop.Det(matop.ScalarMul(cov,2*pi)))
	meanShift := matop.Sub(matop.VecToMatrix(x),matop.VecToMatrix(mean))
	mult := matop.Mult(matop.Mult(matop.Transpose(meanShift), matop.Inv(cov)), meanShift)
	fmt.Println(a, meanShift, mult, matop.Det(mult))
	return (1/a)*math.Exp(-0.5*matop.Det(mult))
}