/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package commandsl

import "math/rand"

var x = 10

func Variables() {
	{
		var y = 20
		println(y)
	}
	var y = 30
	// These sentense has a equally effect of this:
	// y := 30
	// This form can be used in some area where we
	// can not use the normal var keyword.
	// The examble is put below, around 5 lines. A and B.
	println(x, y)
	// The expexted y output is 30 but not 20
	// because the {} declares a domain of each
	// variable.
	if true {
		// TODO: Example A
		for counter := 10; counter > 0; counter-- {
			println(counter)
		}
		// TODO: Example B
		if numb := rand.Intn(3); numb == 0 {
			println("Space Adventures")
		} else if numb == 1 {
			println("China Voyager")
		} else {
			println("Virgin Galactic")
		}
	}
	// aboout transform betweent the domain,
	// we need a thing, which is called pointer.
}
