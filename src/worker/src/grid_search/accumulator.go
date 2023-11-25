package grid_search

import (
	"math"
)

type Accumulator struct {
	callback   Callback
	trueResult float64
	trueInput  [Size]float64
}

type Callback func(float64, [Size]float64)

func NewAccumulator(accumType string) *Accumulator {
	acc := &Accumulator{}
	switch accumType {
	case "MAX":
		acc.callback = acc.max
		acc.trueResult = math.Inf(-1)
	case "MIN":
		acc.callback = acc.min
		acc.trueResult = math.Inf(1)
	default:
		acc.callback = acc.avg
		acc.trueResult = 0
	}
	return acc
}

func (acc *Accumulator) Accumulate(res float64, current [Size]float64) {
	acc.callback(res, current)
}

func (acc *Accumulator) GetResult() float64 {
	return acc.trueResult
}

func (acc *Accumulator) GetInput() [Size]float64 {
	return acc.trueInput
}

func (acc *Accumulator) max(res float64, current [Size]float64) {
	if res > acc.trueResult {
		acc.trueResult = res
		copy(acc.trueInput[:], current[:])
	}
}

func (acc *Accumulator) min(res float64, current [Size]float64) {
	if res < acc.trueResult {
		acc.trueResult = res
		copy(acc.trueInput[:], current[:])
	}
}

func (acc *Accumulator) avg(res float64, current [Size]float64) {
	acc.trueResult += res
}
