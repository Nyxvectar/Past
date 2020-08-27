/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package datac

import "fmt"

func RangeUsage() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	{
		map1 := make(map[int]float32)
		map1[1] = 1.0
		map1[2] = 2.0
		map1[3] = 3.0
		map1[4] = 4.0

		/*
			When a function returns multiple results,
			if you only need the first one, then it
			just only been expected to do XX = FFF
			but if you need the more like the second
			there is in need to write 'XX1, XX2' or
			'_, XX2' and, etc. in same reason.
		*/
		{
			for key, value := range map1 {
				fmt.Printf("key is: %d - value is: %f\n", key, value)
			}
			for key := range map1 {
				fmt.Printf("key is: %d\n", key)
			}
			for _, value := range map1 {
				fmt.Printf("value is: %f\n", value)
			}
		}
	}
}
