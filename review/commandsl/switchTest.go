/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package commandsl

func SwitchTest1() {
	println("There is a cavern entrance here and a path to the east")
	var commanda = "go inside"
	// This usage(1) would directly compare the value they have
	// An expected output is "You find yourself in a dimly lit cavern"
	switch commanda {
	case "go east":
		print("You head further up the hountain")
	case "enter the cave", "go inside":
		println("You find yourself in a dimly lit cavern")
	case "read sign":
		println("The sign reads 'No Minors'.")
	default:
		println("Did not quite get that.")
	}
}

func SwitchTest2() {
	// The switch test 2's form is longer, and if there is no
	// necessary, we prefer to use the form 1.
	var room = "lake"
	switch {
	case room == "cave":
		println("You find yourself in a dimly lit cavern")
	case room == "lake":
		println("The ice seems solid enough.")
		fallthrough
	// The fallthrough can fall through the next branch and run.
	// But we need to know that it would only run the next branch
	// rather than the whole chains of the all case sentences below.
	case room == "underwater":
		println("The ice seems solid enough.(Twice)")
		fallthrough
	case room == "lake":
		println("The ice seems solid enough.(Third)")
	}
}
