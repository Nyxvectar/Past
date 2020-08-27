/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/21/2025
 */

package class

import (
	"fmt"
	"math/rand"
)

func Piggy() {
	var balence = 0.0
	for {
		var t = rand.Intn(3)
		if t == 0 {
			balence = balence + 0.05
		} else if t == 1 {
			balence = balence + 0.10
		} else {
			balence = balence + 0.25
		}
		fmt.Printf("Balence now: %4.2f\n", balence)
		if balence >= 25.00 {
			break
		}
	}
}
