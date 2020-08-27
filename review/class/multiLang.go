/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package class

func TurnToTure(get string) {
	var whether1 = (get == "true")
	var whether2 = (get == "yes")
	var whether3 = (get == "1")
	if whether1 || (whether2 || whether3) == true {
		println("true")
	} else {
		println("pls enter again.")
	}
}

func TurnToFalse(get string) {
	var whether1 = (get == "false")
	var whether2 = (get == "no")
	var whether3 = (get == "0")
	if whether1 || (whether2 || whether3) == true {
		println("false")
	} else {
		println("pls enter again.")
	}
}
