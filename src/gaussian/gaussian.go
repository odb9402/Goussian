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
	"math"
)

const pi = math.Pi

func GaussianUniv(x float64, mean float64, variance float64) float64{
	return 1/math.Sqrt(2*pi*math.Abs(variance))*math.Exp((x - mean)/(2*math.Abs(variance)))
}

//func GaussianMultiv(mean vector, cov matrix) vector{
//
//}