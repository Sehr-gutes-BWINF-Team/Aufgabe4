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
	l := 0
	for i < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 && workTime <= 8 {
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

		if workTime == 24 {
			workTime = 0
		}

		overallTime++
	}

	log.Println(overallTime)
	log.Println(l)
	log.Println("Average", CalculateAverage(finishedOrders), "minutes.")
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
