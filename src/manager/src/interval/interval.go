package interval

import (
	"fmt"
	"math"
)

type Interval struct {
	start     float64
	end       float64
	step      float64
	size      uint64
	precision uint64
}

func NewInterval(start, end, step float64) *Interval {
	size := uint64(math.Ceil((end - start) / step))
	return &Interval{
		start:     start,
		end:       end,
		step:      step,
		size:      size,
		precision: 10,
	}
}

func (interval *Interval) Split(nPartitions uint64) []Interval {

	if nPartitions <= 0 {
		return nil
	}

	if interval.size%nPartitions == 0 {
		return interval.splitEvenly(nPartitions)
	}
	maxElemsPerInterval := int(math.Ceil(float64(interval.size) / float64(nPartitions)))
	nSubIntervalsFull := uint64(math.Floor(float64(interval.size-nPartitions) / float64(maxElemsPerInterval-1)))

	var intervals []Interval
	var subEnd float64
	for j := 0; j < int(nSubIntervalsFull); j++ {
		subStart := roundFloat(interval.start+float64(j*maxElemsPerInterval)*interval.step, float64(interval.precision))
		subEnd = roundFloat(math.Min(interval.end, subStart+float64(maxElemsPerInterval)*interval.step), float64(interval.precision))
		intervals = append(intervals, Interval{start: subStart, end: subEnd, step: interval.step, size: uint64(math.Ceil(subEnd-subStart) / interval.step), precision: interval.precision})
	}
	intervalReminder := NewInterval(subEnd, interval.end, interval.step)
	subIntervalsReminder := intervalReminder.Split(nPartitions - nSubIntervalsFull)
	intervals = append(intervals, subIntervalsReminder...)
	return intervals
}

func (interval *Interval) splitEvenly(nPartitions uint64) []Interval {
	var intervals []Interval
	for j := 0.0; j < float64(nPartitions); j++ {
		subStart := roundFloat(interval.start+j*float64(interval.size)/float64(nPartitions)*interval.step, float64(interval.precision))
		subEnd := roundFloat(interval.start+(j+1)*float64(interval.size)/float64(nPartitions)*interval.step, float64(interval.precision))
		intervals = append(intervals, Interval{start: subStart, end: subEnd, step: interval.step, size: uint64(math.Ceil(subEnd-subStart) / interval.step), precision: interval.precision})
	}
	return intervals
}

func (interval *Interval) print() {
	fmt.Printf("start: %v, end: %v, step: %v\n", interval.start, interval.end, interval.step)
}

func (interval *Interval) intervalSize() uint64 {
	return interval.size
}

func (interval *Interval) getInterval() (float64, float64, float64) {
	return interval.start, interval.end, interval.step
}

func (interval *Interval) setPrecision(precision uint64) {
	interval.precision = precision
}

func roundFloat(value, precision float64) float64 {
	return math.Round(value*math.Pow(10, precision)) / math.Pow(10, precision)
}
