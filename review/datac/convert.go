/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package datac

import (
	"fmt"
	"strconv"
)

func Convert() {
	{
		str := "123"
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error happened:", err)
		} else {
			fmt.Printf("String '%s' convert to int：%d\n", str, num)
		}
	}
}
