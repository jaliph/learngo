https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go


https://stackoverflow.com/questions/39561140/what-is-two-dimensional-arrays-memory-representation

a := make([][]uint8, dy)
for i := range a {
    a[i] = make([]uint8, dx)
}


x := [5][5]byte{}
fmt.Println(&x[0][3])
fmt.Println(&x[0][4])
fmt.Println(&x[1][0])

a := [][]uint8{
    {0, 1, 2, 3},
    {4, 5, 6, 7},
}
fmt.Println(a) // Output is [[0 1 2 3] [4 5 6 7]]


matrix := make([][]int, n)
for i := 0; i < n; i++ {
    matrix[i] = make([]int, m)
}


matrix := make([][]int, n)
rows := make([]int, n*m)
for i := 0; i < n; i++ {
    matrix[i] = rows[i*m : (i+1)*m]
}
// make continous section of memory for faster access