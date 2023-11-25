package interval

import (
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
