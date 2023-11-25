package interval

import (
	"math"
	// "fmt"
)

type Partition struct {
	nPartitions           uint64
	currentPartition      uint64
	nIntervals            int
	iterations            uint64
	intervals             []Interval
	partitionsPerInterval []uint64
	splitIntervals        [][]Interval
	currentIndex          []uint64
}

func NewPartition(intervals []Interval, nIntervals, maxChunkSize int) *Partition {
	p := &Partition{
		nPartitions:           0,
		currentPartition:      0,
		nIntervals:            nIntervals,
		iterations:            0,
		intervals:             intervals,
		partitionsPerInterval: nil,
		splitIntervals:        nil,
		currentIndex:          nil,
	}
	p.Split(maxChunkSize)
	return p
}

func (partition *Partition) Available() bool {
	return partition.currentPartition < partition.nPartitions
}

func (partition *Partition) Next() []Interval {
	_partition := make([]Interval, partition.nIntervals)
	for j := 0; j < partition.nIntervals; j++ {
		_partition[j] = partition.splitIntervals[j][partition.currentIndex[j]]
	}
	for j := 0; j < partition.nIntervals; j++ {
		if partition.currentIndex[j]+1 < partition.partitionsPerInterval[j] {
			partition.currentIndex[j]++
			break
		} else {
			partition.currentIndex[j] = 0
		}
	}
	partition.currentPartition++
	return _partition
}

func (partition *Partition) CalcPartitionPerInterval(minBatches int) []uint64 {
	partitionsPerInterval := make([]uint64, partition.nIntervals)
    for i := range partitionsPerInterval {
        partitionsPerInterval[i] = 1
    }
	var missingPartitions uint64
	var elements uint64
	// fmt.Println("minBatches:", minBatches)
	// fmt.Println("partition.nIntervals:", partition.nIntervals)
	// fmt.Println("partitionsPerInterval:", partitionsPerInterval)
	for i := 0; i < partition.nIntervals; i++ {
		missingPartitions = partition.calcAmountOfMissingPartitions(minBatches, partitionsPerInterval)
		// fmt.Println("missingPartitions:", missingPartitions)
		elements = partition.intervals[i].IntervalSize()
		// fmt.Println("elements:", elements)
		if elements > missingPartitions {
			partitionsPerInterval[i] *= missingPartitions
		} else {
			partitionsPerInterval[i] *= elements
		}
	}
	return partitionsPerInterval
}

func (partition *Partition) calcAmountOfMissingPartitions(minBatches int, partitionsPerInterval []uint64) uint64 {
	// fmt.Println("calcPartitionsAmount:", partition.calcPartitionsAmount(partitionsPerInterval))
	return uint64(math.Ceil(float64(minBatches) / float64(partition.calcPartitionsAmount(partitionsPerInterval))))
}

func (partition *Partition) calcPartitionsAmount(partitionsPerInterval []uint64) uint64 {
	var result uint64 = 1
	// fmt.Println("partitionsPerInterval:", partitionsPerInterval)
	for _, v := range partitionsPerInterval {
		result *= v
	}
	// fmt.Println("result:", result)
	return result
}

func (partition *Partition) Split(maxChunkSize int) {
	// fmt.Println("fullCalculationSize:", partition.fullCalculationSize())

	minBatches := int(math.Floor(float64(partition.fullCalculationSize())/float64(maxChunkSize))) + 1

	partition.partitionsPerInterval = partition.CalcPartitionPerInterval(minBatches)
	// fmt.Println("partitionsPerInterval:", partition.partitionsPerInterval)

	partition.nPartitions = partition.calcPartitionsAmount(partition.partitionsPerInterval)
	// fmt.Println("nPartitions:", partition.calcPartitionsAmount(partition.partitionsPerInterval))
	for i := 0; i < partition.nIntervals; i++ {
		partition.splitIntervals = append(partition.splitIntervals, partition.intervals[i].Split(partition.partitionsPerInterval[i]))
	}
	partition.iterations = partition.calcPartitionsAmount(partition.partitionsPerInterval)
	partition.currentIndex = make([]uint64, partition.nIntervals)
}

func (partition *Partition) fullCalculationSize() uint64 {
	result := uint64(1)
	for i := 0; i < partition.nIntervals; i++ {
		result *= partition.intervals[i].IntervalSize()
		// fmt.Println("result:", partition.intervals[i].IntervalSize())
	}
	return result
}
