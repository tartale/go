package heap

import (
	stdheap "container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_PopsHighestPriorityFirst(t *testing.T) {
	items := []*Item{
		{Value: "low", Priority: 1},
		{Value: "medium", Priority: 2},
		{Value: "high", Priority: 3},
	}

	pq := make(PriorityQueue, 0, len(items))
	stdheap.Init(&pq)
	for _, it := range items {
		stdheap.Push(&pq, it)
	}

	var got []string
	for pq.Len() > 0 {
		item := stdheap.Pop(&pq).(*Item)
		got = append(got, item.Value.(string))
	}

	assert.Equal(t, []string{"high", "medium", "low"}, got)
}

