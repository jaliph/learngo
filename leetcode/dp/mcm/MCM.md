MCM


40, 20, 30, 10, 30
    i            j
    k

temp = cost +  fx(solveleft) + fx(solveRight)
final = min(final, temp)

cost = arr[i-1] * arr[k] * arr[j]

Type 1
loop => k=i; k <= j - 1; k++
solve(i , k), solve(k + 1, j)

calling
solve(1, n - 1)




40, 20, 30, 10, 30
    i           j
                k

Type 2
loop k= i + 1; k <= j; k++
solve(i, k - 1), solve(k, j)


// base 

i >= j
return 0



/// Iterative


40, 20, 30, 10, 30


          i     j