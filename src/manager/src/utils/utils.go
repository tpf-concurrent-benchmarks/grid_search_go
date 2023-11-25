package utils

import (
	"manager/src/interval"
)

type WorkMessage struct {
	Data [][3]float64
	Agg  string
}

func CreateWorkMessageFrom(intervals []interval.Interval, aggregation string) WorkMessage {
	intervalList := make([][3]float64, len(intervals))
	for i, auxInterval := range intervals {
		intervalList[i] = auxInterval.GetInterval()
	}
	message := WorkMessage{
		Data: intervalList,
		Agg:  aggregation,
	}
	return message
}
