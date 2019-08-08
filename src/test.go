package main

import (
	"fmt"
	"gaussian"
	"math/rand"
	"matop"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	var randomFloat [100]float64
	multiFloat := gaussian.RandomMatrix(3, 10000)

	i := 0
	for i < 100 {
		randomFloat[i] = rand.Float64() * 100
		i++
	}

	fmt.Println("EmpMean : ", gaussian.EmpMeanUnivar(randomFloat[:]))
	fmt.Println("EmpVar : ", gaussian.EmpVar(randomFloat[:]))

	fmt.Println("Emp mean Multivar : ", gaussian.EmpMeanMultivar(multiFloat))
	fmt.Println("Emp covariance Multivar : ", gaussian.EmpCovar(multiFloat))
	fmt.Println("Emp correlation coefficient Multivar : ", gaussian.EmpCor(multiFloat))
	fmt.Println("Emp var for 0`th data : ", gaussian.EmpVar(multiFloat[0]))
	fmt.Println("Emp var for 1`th data : ", gaussian.EmpVar(multiFloat[1]))
	fmt.Println("Emp var for 2`th data : ", gaussian.EmpVar(multiFloat[2]))

	m := [][]float64{
		{-1.0, 1.0, -4.0},
		{1.0, -2.0, 0.0},
		{-4.0, 0.0, -3.0},
	}
	fmt.Println("m2 :", matop.Mult(m, m))
	fmt.Println("3m :", matop.ScalarMul(m, 3))
	fmt.Println("mT :", matop.Transpose(m))
	fmt.Println("det(m) :", matop.Det(m))
	fmt.Println("Inv(m) :", matop.Inv(m))
	fmt.Println("m :", m)

	x1 := []float64{
		50, 50, 50,
	}
	x2 := []float64{
		150, 200, -50,
	}

	var n gaussian.MultiGaussian
	n.InitGaussian(gaussian.EmpMeanMultivar(multiFloat), gaussian.EmpCovar(multiFloat))
	fmt.Println("Gaussian prob of x1 data : ", n.GaussianMultiv(x1))
	fmt.Println("Gaussian prob of x2 data : ", n.GaussianMultiv(x2))

	var p gaussian.MultiGaussian
	p.InitGaussian(gaussian.EmpMeanMultivar(multiFloat), m)
	fmt.Println("Gaussian prob of x1 data with another cov: ", p.GaussianMultiv(x1))
	fmt.Println("Gaussian prob of x2 data with another cov: ", p.GaussianMultiv(x2))
	fmt.Println("Gaussian prob of mean data with another cov: ", p.GaussianMultiv(gaussian.EmpMeanMultivar(multiFloat)))

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	plt.Title.Text = "Random data samples"

	scatter_points := make(plotter.XYZs, 100)
	for i := range scatter_points {
		scatter_points[i].X = multiFloat[0][i]
		scatter_points[i].Y = multiFloat[1][i]
		scatter_points[i].Z = multiFloat[2][i]
	}

	scatter_plot, err := plotter.NewScatter(scatter_points)
	if err != nil {
		panic(err)
	}

	plt.Add(scatter_plot)

	if err := plt.Save(4*vg.Inch, 4*vg.Inch, "Test.png"); err != nil {
		panic(err)
	}
}
