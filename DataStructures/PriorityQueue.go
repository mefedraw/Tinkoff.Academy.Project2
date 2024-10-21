package DataStructures

import (
	"TinkofMaze/DataStructures/DataStructuresErrors"
	"TinkofMaze/Maze"
)

type Item struct {
	Point    Maze.Point
	Priority float64
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(item *Item) {
	item.Index = len(*pq)
	*pq = append(*pq, item)
	pq.up(len(*pq) - 1)
}

func (pq *PriorityQueue) Pop() (*Item, error) {
	if pq.Len() == 0 {
		return nil, DataStructuresErrors.NewErrPriorityQueue("Pop from empty priority queue")
	}
	n := len(*pq)
	pq.Swap(0, n-1)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	if pq.Len() > 0 {
		pq.Down(0)
	}
	return item, nil
}

func (pq *PriorityQueue) up(index int) {
	for {
		parent := (index - 1) / 2
		if index == 0 || pq.Less(parent, index) {
			break
		}
		pq.Swap(parent, index)
		index = parent
	}
}

func (pq *PriorityQueue) Down(index int) {
	n := len(*pq)
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		smallest := index
		if leftChild < n && pq.Less(leftChild, smallest) {
			smallest = leftChild
		}
		if rightChild < n && pq.Less(rightChild, smallest) {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		pq.Swap(index, smallest)
		index = smallest
	}
}

func (pq *PriorityQueue) Peek() (*Item, error) {
	if pq.Len() == 0 {
		return nil, DataStructuresErrors.NewErrPriorityQueue("Peek from empty priority queue")
	}
	return (*pq)[0], nil
}

func (pq *PriorityQueue) Update(item *Item, newPriority float64) error {
	if item.Index < 0 || item.Index >= pq.Len() {
		return DataStructuresErrors.NewErrPriorityQueue("Invalid item index for update")
	}
	item.Priority = newPriority
	pq.up(item.Index)
	pq.Down(item.Index)
	return nil
}
