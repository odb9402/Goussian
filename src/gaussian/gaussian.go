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
	"errors"
	"fmt"
	"math"
	"math/rand"
	"matop"
)

const pi = math.Pi

type MultiGaussian struct {
	mean Vector
	cov  Matrix

	cov_inv Matrix // Pre computed inverse of the covariance.
	a       float64
}

func GaussianUniv(x float64, mean float64, variance float64) float64 {
	return 1 / math.Sqrt(2*pi*math.Abs(variance)) * math.Exp((x-mean)/(2*math.Abs(variance)))
}

func (n *MultiGaussian) InitGaussian(mean Vector, cov Matrix) {
	n.mean = mean
	n.cov = cov
	n.cov_inv = matop.Inv(n.cov)
	n.a = math.Sqrt(matop.Det(matop.ScalarMul(n.cov, 2*pi)))
}

func (n MultiGaussian) GaussianMultiv(x Vector) float64 {
	if len(x) != len(n.mean) {
		message := fmt.Sprintf("ERROR: Length of input vector (%d) and mean (%d) are different.\n", len(x), len(n.mean))
		errors.New(message)
		fmt.Println(errors.New(message))
		return 0.0
	}
	meanShift := matop.Sub(matop.VecToMatrix(x), matop.VecToMatrix(n.mean))
	mult := matop.Det(matop.Mult(matop.Mult(matop.Transpose(meanShift), n.cov_inv), meanShift))

	return (1 / n.a) * math.Exp(-0.5*mult)
}

/* It will generate random numbers from a specific gaussian distribution 'n' using
Cholesky decomposition for the covariance matrix and Box-Muller transfrom.
*/
func (n MultiGaussian) Sampling() Vector {
	dim := len(n.cov)
	data := make(Vector, dim)

	for i := 0; i < dim; i++ {
		data[i] = rand.NormFloat64()
	}
	fmt.Println(data)
	fmt.Println(n.cov)
	fmt.Println(matop.VecToMatrix(data))

	fmt.Println(matop.Mult(n.cov, matop.VecToMatrix(data)))

	return data
}
