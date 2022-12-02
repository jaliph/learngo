package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4 // DONOT CHANGE IT!

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
}

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix == 0 {
		return -1
	}
	poppedItem := st.data[st.ix-1]
	st.data[st.ix-1] = 0
	st.ix--
	return poppedItem
}

func (st Stack) String() string {
	var str string
	for i := 0; i < st.ix; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(st.data[i]) + "]"
	}
	return str
}
