// https://leetcode.com/problems/find-all-people-with-secret/

package graph

import (
	"fmt"
	"sort"
)

// Refer dijkstra
func findAllPeople_BFS(n int, meetings [][]int, firstPerson int) []int {

	// sort based on the meetings
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	seenTime := map[int]struct{}{}
	sortedMeetingsGroup := [][][]int{}

	for _, meeting := range meetings {
		if _, ok := seenTime[meeting[2]]; !ok {
			sortedMeetingsGroup = append(sortedMeetingsGroup, [][]int{})
		}
		seenTime[meeting[2]] = struct{}{}
		idx := len(sortedMeetingsGroup) - 1
		sortedMeetingsGroup[idx] = append(sortedMeetingsGroup[idx], meeting)
	}

	knowSecret := map[int]struct{}{
		0:           {},
		firstPerson: {},
	}

	for _, meetingGroup := range sortedMeetingsGroup {
		peopleknowSecret := map[int]struct{}{}
		meetingGraph := map[int][]int{}

		// setup the graph for each node
		for _, person := range meetingGroup {
			if _, ok := meetingGraph[person[0]]; !ok {
				meetingGraph[person[0]] = []int{}
			}
			if _, ok := meetingGraph[person[1]]; !ok {
				meetingGraph[person[1]] = []int{}
			}
			meetingGraph[person[0]] = append(meetingGraph[person[0]], person[1])
			meetingGraph[person[1]] = append(meetingGraph[person[1]], person[0])

			if _, ok := knowSecret[person[0]]; ok {
				peopleknowSecret[person[0]] = struct{}{}
			}

			if _, ok := knowSecret[person[1]]; ok {
				peopleknowSecret[person[1]] = struct{}{}
			}
		}

		queue := []int{}
		for person := range peopleknowSecret {
			queue = append(queue, person)
		}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			knowSecret[curr] = struct{}{}

			for _, neigh := range meetingGraph[curr] {
				if _, ok := knowSecret[neigh]; !ok {
					knowSecret[neigh] = struct{}{}
					queue = append(queue, neigh)
				}
			}
		}
	}

	list := []int{}

	for person := range knowSecret {
		list = append(list, person)
	}
	return list
}

func Driver() {
	n := 3
	meetings := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson := 3
	fmt.Println(findAllPeople_BFS(n, meetings, firstPerson))
}
