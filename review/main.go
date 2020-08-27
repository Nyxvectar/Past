/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package main

import (
	"fmt"
	"siracusan/block"
	"siracusan/class"
	"siracusan/commandsl"
	"siracusan/datac"
)

func main() {
	initation(false)
	switchCmd(false)
	switchClass(false)
	DraftEnable(false)
	block.InterfaceTrain()
}

func switchClass(whether bool) {
	if whether == true {
		{
			class.Piggy()
			class.Consts()
			class.TurnToTure("1")
			class.TurnToFalse("1")
		}
		{
			block.Demo(10, 20)
			block.Struc()
			block.PointerTrain()
			block.Book()
		}
		{
			datac.Slices()
			datac.RangeUsage()
			datac.MapT()
			datac.CaseTrain()
			datac.Convert()
		}
	}
}

func switchCmd(whether bool) {
	if whether == true {
		commandsl.CalculateG()
		commandsl.ForWhile()
		commandsl.LogicalT()
		commandsl.YearCheck(2002)
		commandsl.SwitchTest2()
		commandsl.ForL(10)
		commandsl.Infinity(10)
		commandsl.Variables()
		commandsl.Randomly(50)
	}
}

func initation(whether bool) {
	if whether == true {
		s := "gopher"
		fmt.Printf("Hello and welcome, %s!\n", s)
		for i := 1; i <= 5; i++ {
			fmt.Println("i =", 100/i)
		}
	}
}
