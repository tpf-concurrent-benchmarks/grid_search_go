package grid_search

import (
	"math"
)

func griewankFun(parameters [3]float64) float64 {
	a := parameters[0]
	b := parameters[1]
	c := parameters[2]
	return (1.0/4000.0)*(a*a+b*b+c*c) - math.Cos(a)*math.Cos(b/math.Sqrt(2))*math.Cos(c/math.Sqrt(3)) + 1
}
