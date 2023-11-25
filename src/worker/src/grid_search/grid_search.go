package grid_search

type GridSearch struct {
	params      *Params
	result      float64
	totalInputs int
	accumType   string
	input       [Size]float64
}

func NewGridSearch(params *Params, accumType string) *GridSearch {
	return &GridSearch{
		params:      params,
		result:      0.0,
		totalInputs: 0,
		accumType:   accumType,
		input:       [Size]float64{},
	}
}

func (gs *GridSearch) search(callback func([Size]float64) float64) {
	accumulator := NewAccumulator(gs.accumType)

	for i := int64(0); i < gs.params.getTotalIterations(); i++ {
		current := gs.params.getCurrent()
		res := callback(current)
		accumulator.accumulate(res, current)
		gs.params.next()
	}
	gs.result = accumulator.getResult()
	gs.input = accumulator.getInput()
}

func (gs *GridSearch) getTotalInputs() int {
	return gs.totalInputs
}

func (gs *GridSearch) getResult() float64 {
	return gs.result
}

func (gs *GridSearch) getInput() [Size]float64 {
	return gs.input
}
