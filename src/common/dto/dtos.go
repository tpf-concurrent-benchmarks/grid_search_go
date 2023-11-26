package dto

type WorkMessage struct {
	Data [][3]float64 `json:"data"`
	Agg  string       `json:"agg"`
}

type MaxResultsDTO struct {
	Value      float64    `json:"value"`
	Parameters [3]float64 `json:"parameters"`
}

type MinResultsDTO struct {
	Value      float64    `json:"value"`
	Parameters [3]float64 `json:"parameters"`
}

type AvgResultsDTO struct {
	Value            float64 `json:"value"`
	ParametersAmount uint64  `json:"paramsAmount"`
}
