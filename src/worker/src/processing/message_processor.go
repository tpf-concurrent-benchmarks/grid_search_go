package processing

import (
	"shared/dto"
	"worker/src/grid_search"
)

func ProcessMessage(message *dto.WorkMessage) {
	aggregation := message.Agg
	parameters := message.Data
	start, end, step := [grid_search.Size]float64{}, [grid_search.Size]float64{}, [grid_search.Size]float64{}
	for i := 0; i < len(parameters); i++ {
		start[i] = parameters[i][0]
		end[i] = parameters[i][1]
		step[i] = parameters[i][2]
	}
	params := grid_search.NewParams(start, end, step)
	gridSearch := grid_search.NewGridSearch(params, aggregation)
	gridSearch.Search(grid_search.GriewankFunc)

	results := aggregate(gridSearch, aggregation)
	// TODO send results to manager, look up how to get the conn client
}

func aggregate(gridSearch *grid_search.GridSearch, aggregation string) interface{} {
	// TODO complete this function
	switch aggregation {
	case "MAX":
		return sum(gridSearch)
	case "MIN":
		return min(gridSearch)
	case "AVG":
		return avg(gridSearch)
	}
	return nil
}
