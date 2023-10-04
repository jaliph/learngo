// https://leetcode.com/problems/detect-squares/

package maths

import "fmt"

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type DetectSquares struct {
	list []Point
	pMap map[string]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		list: []Point{},
		pMap: make(map[string]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	p := NewPoint(point[0], point[1])
	this.list = append(this.list, p)
	this.pMap[Key(point[0], point[1])]++
}

func (this *DetectSquares) Count(point []int) int {
	count := 0
	for _, p := range this.list {
		if (abs(point[0]-p.x)) == (abs(point[1]-p.y)) && point[0] != p.x && point[1] != p.y {
			v1, ok1 := this.pMap[Key(point[0], p.y)]
			v2, ok2 := this.pMap[Key(p.x, point[1])]
			if ok1 && ok2 {
				count += (v1 * v2)
			}
		}
	}
	return count
}

func Key(x, y int) string {
	return fmt.Sprintf("%d#%d", x, y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
