package interval

import (
	"math"
	"testing"
)

func TestNewPartition(t *testing.T) {
	intervals := []Interval{
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	}
	nIntervals := len(intervals)
	maxChunkSize := 5

	partition := NewPartition(intervals, nIntervals, maxChunkSize)

	if partition.nIntervals != nIntervals {
		t.Errorf("Expected nIntervals to be %d, got %d", nIntervals, partition.nIntervals)
	}

	if len(partition.intervals) != nIntervals {
		t.Errorf("Expected intervals length to be %d, got %d", nIntervals, len(partition.intervals))
	}
}

func TestPartitionsPerInterval(t *testing.T) {
	intervals := []Interval{
		*NewInterval(4, 8, 1),
		*NewInterval(4, 8, 1),
		*NewInterval(8, 10, 1),
	}
	nIntervals := len(intervals)
	maxChunkSize := 5

	partition := NewPartition(intervals, nIntervals, maxChunkSize)

	partitionsPerInterval := partition.CalcPartitionPerInterval(3)

	t.Logf("Partitions per interval: %v", partitionsPerInterval)
	if len(partitionsPerInterval) != nIntervals {
		t.Errorf("Expected partitionsPerInterval length to be %d, got %d", nIntervals, len(partitionsPerInterval))
	}

	expected := []uint64{3, 1, 1}

	for i := 0; i < nIntervals; i++ {
		if partitionsPerInterval[i] != expected[i] {
			t.Errorf("Expected partitionsPerInterval[%d] to be %d, got %d", i, expected[i], partitionsPerInterval[i])
		}
	}
}

func TestAvailable(t *testing.T) {
	intervals := []Interval{
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	}
	nIntervals := len(intervals)
	maxChunkSize := 5

	partition := NewPartition(intervals, nIntervals, maxChunkSize)

	available := partition.Available()

	if !available {
		t.Errorf("Expected Available to return true, got false")
	}
}

func TestAllAvailable(t *testing.T) {
	intervals := []Interval{
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	}
	nIntervals := len(intervals)
	maxChunkSize := 5

	partition := NewPartition(intervals, nIntervals, maxChunkSize)
	minBatches := int(math.Floor(float64(partition.fullCalculationSize())/float64(maxChunkSize))) + 1

	for i := 0; i < minBatches; i++ {
		available := partition.Available()

		if !available {
			t.Errorf("Expected Available to return true, got false")
		}
		partition.Next()
	}
}

func TestPartitionsOne(t *testing.T) {
	intervals := []Interval{
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	}
	nIntervals := len(intervals)
	maxChunkSize := 5

	partition := NewPartition(intervals, nIntervals, maxChunkSize)

	expected := [][]Interval{{
		*NewInterval(0, 5, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	},
		{
			*NewInterval(5, 10, 5),
			*NewInterval(0, 10, 5),
			*NewInterval(0, 10, 5),
		},
	}

	i := 0
	for partition.Available() {
		part := partition.Next()
		for j := 0; j < len(part); j++ {
			if !compareIntervals(&part[j], &expected[i][j]) {
				t.Errorf("Expected %v, got %v", expected[i][j], part[j])
			}
		}
		i++
	}
}

func TestPartitionsMultiple(t *testing.T) {
	intervals := []Interval{
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
		*NewInterval(0, 10, 5),
	}
	nIntervals := len(intervals)
	maxChunkSize := 4

	partition := NewPartition(intervals, nIntervals, maxChunkSize)

	expected := [][]Interval{{
		*NewInterval(0, 5, 5),
		*NewInterval(0, 5, 5),
		*NewInterval(0, 10, 5),
	},
		{
			*NewInterval(5, 10, 5),
			*NewInterval(0, 5, 5),
			*NewInterval(0, 10, 5),
		},
		{
			*NewInterval(0, 5, 5),
			*NewInterval(5, 10, 5),
			*NewInterval(0, 10, 5),
		},
		{
			*NewInterval(5, 10, 5),
			*NewInterval(5, 10, 5),
			*NewInterval(0, 10, 5),
		},
	}

	i := 0
	for partition.Available() {
		part := partition.Next()
		for j := 0; j < len(part); j++ {
			if !compareIntervals(&part[j], &expected[i][j]) {
				t.Errorf("Expected %v, got %v", expected[i][j], part[j])
			}
		}
		i++
	}
}
