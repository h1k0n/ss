package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground", topKFrequent([]int{1}, 1))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	pq := make(priorityqueue, 0)
	for v, c := range m {
		e := elem{
			value: v,
			count: c,
		}
		heap.Push(&pq, e)
		if len(pq) > k {
			heap.Pop(&pq)
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		x := heap.Pop(&pq)
		x1 := x.(elem)
		res[i] = x1.value
	}
	return res
}

type elem struct {
	value int
	count int
}
type priorityqueue []elem

func (pq priorityqueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}
func (pq priorityqueue) Len() int {
	return len(pq)
}
func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityqueue) Push(e interface{}) {
	*pq = append(*pq, e.(elem))
}
func (pq *priorityqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x

}
