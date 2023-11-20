package grid_search

import (
	"math"
)

const Size = 3

type Params struct {
	start           [Size]float64
	end             [Size]float64
	step            [Size]float64
	current         [Size]float64
	totalIterations int64
}

func NewParams(start, end, step [Size]float64) *Params {
	params := &Params{
		start:           start,
		end:             end,
		step:            step,
		current:         start,
		totalIterations: 1,
	}
	for i := 0; i < Size; i++ {
		cumParam := int64(math.Floor((end[i] - start[i]) / step[i]))
		if cumParam == 0 {
			cumParam = 1
		}
		params.totalIterations *= cumParam
	}
	return params
}

func (params *Params) getCurrent() [Size]float64 {
	return params.current
}

func (params *Params) next() {
	for i := Size - 1; i >= 0; i-- {
		if params.current[i]+params.step[i] < params.end[i] {
			params.current[i] += params.step[i]
			break
		} else {
			params.current[i] = params.start[i]
		}
	}
}

func (params *Params) getTotalIterations() int64 {
	return params.totalIterations
}
