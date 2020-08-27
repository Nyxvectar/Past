/**
 * Author:  Nyxvectar Yan
 * Repo:    gocy
 * Created: 07/22/2025
 */

package main

import "fmt"

func DraftEnable(enable bool) {
	if enable {
		Draft()
	}
}

func Draft() {
	var a = 21
	var b = 10
	{
		if a == b {
			fmt.Printf("a = b\n")
		} else {
			fmt.Printf("a != b\n")
		}
		a = 5
		b = 20
		if a <= b {
			fmt.Printf("a <= b\n")
		}
		if b >= a {
			fmt.Printf("b >= a\n")
		}
	}
	{
		if a < b {
			fmt.Printf("a < 2\n")
		} else {
			fmt.Printf("a !< b\n")
		}
		if a > b {
			fmt.Printf("a > b\n")
		} else {
			fmt.Printf("a !> b\n")
		}
	}
}
