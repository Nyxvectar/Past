/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package block

import "fmt"

func Struc() {
	var i, n = 0, 0
	{
		// This is a array.
		// The length of it can not be edited.
		// an expected output is '2'
		var arrayName = [2]int{1, 2}
		temp := arrayName[1]
		fmt.Println(temp)
		{
			var Ett [10]int
			for i < 10 {
				Ett[n] = i
				fmt.Printf("Element[%v] = %v\n", n, Ett[i])
				n++
				i++
			}
			// A classical manner to print a array is like this.
			//  for k = 0; k < 'TheLengthYouKnow'; k++ {
			//      fmt.Printf("balance3[%d] = %f\n", k, arrayName[k] )
			//  }
			{
				var j, k int
				balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
				{
					for j = 0; j < 5; j++ {
						fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
					}
					balance3 := [5]float32{1: 2.0, 3: 7.0}
					for k = 0; k < 5; k++ {
						fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
					}
				}
			}
		}
	}
}
