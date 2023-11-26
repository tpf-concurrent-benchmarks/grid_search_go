package grid_search

type GridSearch struct {
	params      *Params
	result      float64
	totalInputs uint64
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

func (gs *GridSearch) Search(callback func([Size]float64) float64) {
	accumulator := NewAccumulator(gs.accumType)

	for i := int64(0); i < gs.params.getTotalIterations(); i++ {
		current := gs.params.getCurrent()
		res := callback(current)
		accumulator.Accumulate(res, current)
		gs.params.next()
	}
	gs.result = accumulator.GetResult()
	gs.input = accumulator.GetInput()
}

func (gs *GridSearch) GetTotalInputs() uint64 {
	return gs.totalInputs
}

func (gs *GridSearch) GetResults() float64 {
	return gs.result
}

func (gs *GridSearch) GetInput() [Size]float64 {
	return gs.input
}
