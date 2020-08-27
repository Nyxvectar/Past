/**
 * Author:  Nyxvectar Yan
 * Repo:    siracusan
 * Created: 07/22/2025
 */

package datac

// Map usage:
// map_variable := make(map[KeyType]ValueType, initialCapacity)

func MapT() {
	var m = map[string]string{
		"china": "UTC+8",
		"usa":   "UTC-7",
		"uk":    "UTC",
	}
	print(m["china"])
	var values, ok = m["yan"]
	if ok != true {
		println("\nWe do not know the time block of Yan")
		println("The length of the Map is", len(m))
	} else {
		println("The time block of Yan is", values)
	}
}
