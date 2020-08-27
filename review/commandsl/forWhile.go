/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package commandsl

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var walkOutside = true
var takeTheBlueHill = false
var command = "Walk outside"

func ForWhile() {
	boolTest(true)
	elseIfTest(1)
}

func elseIfTest(order int) {
	//The variables name is in the front and the class
	//of it should be put behind its name.
	//var judgeResult = strings.Contains(order, "")
	if order >= 0 {
		//If func can be put into a if, and it is same
		//for the other case like while, case, etc.
		if order <= 100 {
			if order == 1 {
				println("The num is 1")
			} else if order == 2 {
				println("The num is 2")
			} else {
				println("The num is %v", order)
			}
			//Just for fun, do not ask why not directly get %v
		}
	} else {
		print("exit")
		os.Exit(0)
	}
}

func boolTest(whether bool) {
	fmt.Println("You found that you was in a cavern")
	var exit = strings.Contains(command, "outside")
	//The check is sensitive with upper or lower case.
	//Contains can check the "" rather than "outside"
	//It seems that it means a form of boolean true.
	fmt.Println("You left the cave:", exit)
	if whether == true {
		if walkOutside == true {
			if takeTheBlueHill == true {
				print("Got it.")
			} else {
				print("Failed")
			}
		}
	}
}

// a possible list of usually syns are
// > < >= <= == != etc.

func ForL(count int) {
	for count > 0 {
		println(count)
		time.Sleep(time.Second)
		count--
	}
}

func Infinity(degrees int) {
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 20 {
			degrees = 0
			// each time degrees got 20 then it would trigger once
			// random, so if it broke, the last num would be 20 forever.
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
	println("If this sentence appears, golang still running after the break in for if.")
}
