package matrix

import (
	"math"
)

type matrix [][]float64
type vector []float64

func Mult(x matrix, y matrix) [][]float64{
	if (len(x[0]) != len(y)) {
		return nil
	} else {
		z := make(matrix, len(x))
		for i := 0; i < len(x); i++{
			z[i] = make(vector, len(y[0]))
		}
		
		for i := 0; i < len(x); i++{
			for j := 0; j < len(x) ; j++{
				z[i][j] = x[i][j] * y[j][i]
			}
		}
		return z
	}
}

func Add(x matrix, y matrix) [][]float64{
	if (len(x) != len(y) || len(x[0]) != len(y[0])){
		return nil
	} else {
		for i := 0; i < len(x); i++{
			for j := 0; j < len(x) ; j++{
				x[i][j] = x[i][j] + y[i][j]
			}
		}
		return x
	}
}

func ScalarMul(x matrix, a float64) [][]float64{
	for i := 0; i < len(x); i++{
		for j := 0; j < len(x[0]); j++{
			x[i][j] = a * x[i][j]
		}
	}
	return x
}

func Inv(x matrix) [][]float64{

}

func Det(x matrix) [][]float64{

}