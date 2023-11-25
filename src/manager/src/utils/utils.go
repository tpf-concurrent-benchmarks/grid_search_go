package utils

import (
	"manager/src/interval"
	"shared/dto"
)

func CreateWorkMessageFrom(intervals []interval.Interval, aggregation string) dto.WorkMessage {
	intervalList := make([][3]float64, len(intervals))
	for i, auxInterval := range intervals {
		intervalList[i] = auxInterval.GetInterval()
	}
	message := dto.WorkMessage{
		Data: intervalList,
		Agg:  aggregation,
	}
	return message
}
