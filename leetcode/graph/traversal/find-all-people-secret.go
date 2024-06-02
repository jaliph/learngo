// https://leetcode.com/problems/find-all-people-with-secret/

package graph

import (
	"container/heap"
	"fmt"
)

type Meeting struct {
	time   int
	person int
}

type PriorityQueue []*Meeting

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	meeting := x.(*Meeting)
	*pq = append(*pq, &Meeting{
		time:   meeting.time,
		person: meeting.person,
	})
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {

	graph := map[int][][]int{}
	for _, meeting := range meetings {
		if _, ok := graph[meeting[0]]; !ok {
			graph[meeting[0]] = [][]int{}
		}
		graph[meeting[0]] = append(graph[meeting[0]], []int{meeting[1], meeting[2]})

		if _, ok := graph[meeting[1]]; !ok {
			graph[meeting[1]] = [][]int{}
		}
		graph[meeting[1]] = append(graph[meeting[1]], []int{meeting[0], meeting[2]})
	}

	visited := make([]bool, n)
	queue := PriorityQueue{}

	queue = append(queue, &Meeting{
		time:   0,
		person: 0,
	})
	queue = append(queue, &Meeting{
		time:   0,
		person: firstPerson,
	})
	heap.Init(&queue)

	for len(queue) > 0 {
		curr := heap.Pop(&queue).(*Meeting)

		if visited[curr.person] {
			continue
		}
		visited[curr.person] = true

		for _, neigh := range graph[curr.person] {
			if !visited[neigh[0]] && curr.time <= neigh[1] {
				heap.Push(&queue, &Meeting{
					time:   neigh[1],
					person: neigh[0],
				})
			}
		}
	}

	people := []int{}
	for person, value := range visited {
		if value {
			people = append(people, person)
		}
	}

	return people
}

func Driver() {
	n := 5
	meetings := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson := 3
	fmt.Println(findAllPeople(n, meetings, firstPerson))
}
