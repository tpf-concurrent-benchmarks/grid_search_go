package dto

type WorkMessage struct {
	Data [][3]float64 `json:"data"`
	Agg  string       `json:"agg"`
}
