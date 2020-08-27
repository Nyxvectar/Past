/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package datac

import "fmt"

func Slices() {
	var s = make([]int, 5, 8)
	var u = s
	{
		printSlice(s)
		s = append(s, 1, 2, 3, 4, 5)
		printSlice(s)
	}
	/*
		The slice is using a same basic, so if
		You changed the variable u, the s will
		be edited together, due to they are us
		ing a same basical array.
	*/
	u[1] = 2
	printSlice(u)
	/*
		This is an advantage of go slice, due
		to its improvement of performance and
		memory usage.
	*/
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
