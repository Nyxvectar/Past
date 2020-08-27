/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package datac

import "fmt"

func CaseTrain() {
	var i int
	for i = 0; i < 10000; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	println(factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
