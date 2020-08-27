/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package commandsl

var aa = true
var bb = false

func LogicalT() {
	if aa || bb {
		// Expected is true because there had a true
		// within aa and bb
		print("aa || bb is ", aa || bb)
	} else if aa && bb {
		print("aa && bb is ", aa && bb)
	} else {
		print("both || and && could not run")
	}
}

func YearCheck(year int) {
	if year%100 == 0 {
		if year%400 == 0 {
			print("yes")
		} else {
			print("nope")
		}
	} else if year%4 == 0 {
		print("yes")
	} else {
		print("nope")
	}
}
