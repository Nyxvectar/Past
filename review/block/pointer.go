/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package block

import "fmt"

func PointerTrain() {
	var NilP *int
	var temp1, temp2 = 1, 2
	var result1, result2 *int
	// '&' at here means 'Get the address'
	// '*' at here means 'Get the Object'
	result1, result2 = &temp1, &temp2
	var a, b = *result1, *result2
	{
		println("The plus answer is", a+b)
		println("Nil output is like:", NilP)
		// The nil one expected a output like 0x0
	}
	{
		if NilP == nil {
			println("NilP is a nil pointer")
		} else if NilP != nil {
			println("NilP is not a nil pointer")
		} else {
			print("Something wrong happened")
		}
	}
	/*
		Note:
		A affective way to use {} is like this:
		for i := 0; i < 1e6; i++ {
			 {
			 	 temp := newBuffer()
			  	 process(temp)
			 }
		}
		temp will be recycled immediately after
		the previous for syntax done.
		An improper way to use {} is like this:
	*/
	{
		var x = 1
		println(x)
		/*
			The memory benefits is too low,
			So there is no need to use {}.
		*/
		{

			// But if you just want to made the code
			// easy to be read, then for sure that
			// the {} is a good choice to select.
			var tt int
			var ptr *int
			var pptr **int
			tt = 3000
			ptr = &tt
			pptr = &ptr
			fmt.Printf("a = %d\n", tt)
			fmt.Printf("Pointer *ptr = %d\n", *ptr)
			fmt.Printf("Pointer to Pointer **pptr = %d\n", **pptr)
		}
	}
}
