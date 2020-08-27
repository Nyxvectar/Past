// 2024-12-21
// Go 1.23

package main

import (
	"fmt"
	"sort"
)

type Station struct {
	distance float64
	price    float64
}

func main() {
	var d1, c, d2, p float64
	var n int
	_, err := fmt.Scan(&d1, &c, &d2, &p, &n)
	if err != nil {
		return
	}

	stations := make([]Station, n+1)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&stations[i].distance, &stations[i].price)
		if err != nil {
			return
		}
	}
	stations[n] = Station{distance: d1, price: 0}

	sort.Slice(stations, func(i, j int) bool {
		return stations[i].distance < stations[j].distance
	})

	if n > 0 && stations[0].distance != 0 {
		fmt.Println("No Solution")
		return
	}

	currentPosition := 0.0
	remainFuel := 0.0
	totalCost := 0.0
	currentStationIndex := 0

	for currentPosition < d1 {
		maxDistance := currentPosition + c*d2
		if currentStationIndex+1 >= len(stations) || stations[currentStationIndex+1].distance > maxDistance {
			if maxDistance >= d1 {
				needFuel := (d1 - currentPosition) / d2
				if needFuel > remainFuel {
					totalCost += (needFuel - remainFuel) * p
				}
				break
			} else {
				fmt.Println("No Solution")
				return
			}
		}

		nextStationIndex := -1
		cheaperFound := false
		for i := currentStationIndex + 1; i < len(stations); i++ {
			if stations[i].distance > maxDistance {
				break
			}
			if stations[i].price < p {
				nextStationIndex = i
				cheaperFound = true
				break
			}
			if nextStationIndex == -1 || stations[i].price < stations[nextStationIndex].price {
				nextStationIndex = i
			}
		}

		if !cheaperFound {
			needFuelToNext := (stations[nextStationIndex].distance - currentPosition) / d2
			fuelToFill := c - remainFuel
			totalCost += fuelToFill * p
			remainFuel = c - needFuelToNext
			currentPosition = stations[nextStationIndex].distance
			p = stations[nextStationIndex].price
			currentStationIndex = nextStationIndex
		} else {
			needFuelToNext := (stations[nextStationIndex].distance - currentPosition) / d2
			if needFuelToNext > remainFuel {
				totalCost += (needFuelToNext - remainFuel) * p
				remainFuel = 0
			} else {
				remainFuel -= needFuelToNext
			}
			currentPosition = stations[nextStationIndex].distance
			p = stations[nextStationIndex].price
			currentStationIndex = nextStationIndex
		}
	}
	fmt.Printf("%.2f\n", totalCost)
}
