package matop

import (
	"math"
	"fmt"
	"errors"
)

type Matrix [][]float64
type Vector []float64

/*
*	Note that intent of all these Matrix operations in this source
*	does not change original operand matrixes by copying.
*/

func VecToMatrix(x Vector) [][]float64{
	/*
	* Convert Vector to n*1 size Matrix.
	*/
	m := make(Matrix, len(x))
	for i:=0; i<len(m); i++{
		m[i] = make(Vector, 1)
		m[i][0] = x[i]
	}
	return m
}

func Mult(x Matrix, y Matrix) [][]float64{
	if (len(x[0]) != len(y)) {
		return nil
	} else {
		z := make(Matrix, len(x))
		for i := 0; i < len(x); i++{
			z[i] = make(Vector, len(y[0]))
		}
		
		for i := 0; i < len(x); i++{
			for j := 0; j < len(x) ; j++{
				z[i][j] = x[i][j] * y[j][i]
			}
		}
		return z
	}
}

func Add(x Matrix, y Matrix) [][]float64{
	if (len(x) != len(y) || len(x[0]) != len(y[0])){
		message := fmt.Sprintf("ERROR: Size of matrix (%d,%d),(%d,%d) are different.\n", len(x),len(x[0]),len(y),len(y[0]))
		fmt.Println(errors.New(message))
		return nil
	} else {
		z := make(Matrix, len(x))
		for i := 0; i < len(x); i++{
			z[i] = make(Vector, len(y[0]))
		}
		for i := 0; i < len(x); i++{
			for j := 0; j < len(x[i]) ; j++{
				z[i][j] = x[i][j] + y[i][j]
			}
		}
		return z 
	}
}

func ScalarMul(x Matrix, a float64) [][]float64{
	z := make(Matrix, len(x))
	for i := 0; i < len(x); i++{
		z[i] = make(Vector, len(x[0]))
	}
	for i := 0; i < len(x); i++{
		for j := 0; j < len(x[0]); j++{
			z[i][j] = a * x[i][j]
		}
	}
	return z
}

func Minus(x Matrix) [][]float64{
	return ScalarMul(x,-1)
}

func Sub(x Matrix, y Matrix) [][]float64{
	return Add(x, Minus(y))
}

func Transpose(x Matrix) [][]float64{
	x_T := make(Matrix, len(x[0]))
	for i := 0; i < len(x[0]); i++{
		x_T[i] = make(Vector, len(x))
	}

	for i := 0; i < len(x[0]); i++{
		for j := 0; j < len(x); j++{
			x_T[i][j] = x[j][i]
		}
	}
	return x_T
}

func luDecomp(x Matrix) [][]float64{
	/*
	* Calculate LU decomposition of a Matrix x.
	* For memory efficiency, the result of decomposition 
	* is stored in a single Matrix LU that includes L in
	* Lower triangle in the Matrix and U in Upper triangle.
	*/
	
	LU := make(Matrix, len(x))
	for i := 0; i < len(x[0]); i++{
		LU[i] = make(Vector, len(x[0]))
	}
	smallDim := int(math.Min(float64(len(x)), float64(len(x[0]))))
	a := 0.0

	//Copy x into LU
	for i := 0; i < len(x); i++{
		for j := 0; j < len(x[i]); j++{
			LU[i][j] = x[i][j]
		}
	}

	for i := 0; i < smallDim; i++{
		a = 1.0 / LU[i][i]
		for j := i+1; j < len(x); j++{
			LU[j][i] = LU[j][i] * a
		}
		for j := i+1; j < len(x[0]); j++{
			for k := i+1; k < len(x); k++{
				LU[j][k] = LU[j][k] - LU[j][i]*LU[i][k]
			}
		}
	}
	return LU
}

func Det(x Matrix) float64{
	/*
	* Calculate determinant of a Matrix.
	* Note that det det(A) = det(LU) = det(L)*det(U) and
	* a triangle Matrix has determinant of sum of product of
	* diagonal elements.
	*/
	det := 1.0
	if len(x) == 1{
		return x[0][0]
	} else {
		LU := luDecomp(x)
		for i := 0; i < len(x); i++{
			det *= LU[i][i]
		}
		return det
	}
}

func equationSol(x Matrix, y Vector) []float64{
	LU := luDecomp(x)
	sol := make(Vector, len(y))
	for i := 0; i < len(sol); i++{
		sol[i] = y[i]
	}

	// solution for L
	for i := 0; i < len(sol); i++{
		for j := 0; j < i; j++{
			sol[i] = sol[i] - LU[i][j]*sol[j]
		}
	}

	// solution for U
	for k := 0; k < len(sol); k++{
		i := len(y) - (k+1)
		for j := i+1; j < len(y) ; j++{
			sol[i] = sol[i] - LU[i][j]*sol[j]
		}
		sol[i] = sol[i] / LU[i][i]
	}
	return sol 
}

func Inv(x Matrix) [][]float64{
	/*
	* Calculate Matrix inversion using LUP decomposition.
	* 1. decompose the to-be-inverted Matrix as lower(L) and
	* upper(U) triangular matrices.
	* 2. record its interchanges in Matrix P
	* 
	* PA = LU
	* PAx = Pb
	* LUx = Pb
	*/
	xInv := make(Matrix, len(x))
	for i := 0 ; i < len(x); i++{
		xInv[i] = make(Vector, len(x[i]))
	}

	for i := 0; i < len(xInv); i++{
		e := make(Vector, len(x[i]))
		for j := range e{
			e[j] = 0
		}
		e[i] = 1
		xInv[i] = equationSol(x, e)
	}
	return xInv
}