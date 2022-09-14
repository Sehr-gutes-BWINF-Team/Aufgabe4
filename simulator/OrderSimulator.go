package simulator

import (
	"log"
)

// Simulate simulates the bike repair shop scenario.
func Simulate(simulationLength int, orders []Order) {
	log.Println("Started simulation!")

	var currentOrders []Order
	var finishedOrders []int
	i, overallTime, workTime := 0, 0, 0
	for i < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 && workTime <= 7 {
			currentOrder = &currentOrders[0]
			currentOrder.workOn()

			if currentOrder.isCompleted() {
				waitingTime := overallTime - currentOrder.Entry
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = Remove(currentOrders, 0)
				i++
			}
		}

		workTime++

		if workTime == 23 {
			workTime = 0
		}

		overallTime++
	}

	average := CalculateAverage(finishedOrders)
	log.Println("Average:", average, "minutes.")
	log.Println("Which are equal to", average/60, "hours.")
	log.Println("Or", average/60/24, "days.")
	log.Println("Stopped simulation!")
}

func getNewOrders(now int, orders []Order) []Order {
	var newOrders []Order
	for i := range orders {
		if now == orders[i].Entry {
			newOrders = append(newOrders, orders[i])
		}
	}
	return newOrders
}
