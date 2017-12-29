package helper

// Permut returns all possible permutations from an array of int
func Permut(arr []int) [][]int {
	var helper func(int, []int)
	res := [][]int{}

	helper = func(n int, arr []int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(n-1, arr)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(len(arr), arr)
	return res
}

/*
https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go

Here's code that iterates over all permutations without generating them all first. The slice p keeps the intermediate state as offsets in a Fisher-Yates shuffle algorithm. This has the nice property that the zero value for p describes the identity permutation.

package main

import "fmt"

func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm(orig, p []int) []int {
    result := append([]int{}, orig...)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}

func main() {
    orig := []int{11, 22, 33}
    for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
        fmt.Println(getPerm(orig, p))
    }
}

*/
