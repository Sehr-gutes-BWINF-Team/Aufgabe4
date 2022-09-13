package simulator

import (
	"log"
)

// Simulate simulates the bike repair shop scenario.
func Simulate(simulationLength int, orders []Order) {
	log.Println("Started simulation!")

	var currentOrders []Order
	var finishedOrders []int
	i := 0
	now := 0
	for i < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(now, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 {
			currentOrder = &currentOrders[0]
			currentOrder.workOn()

			if currentOrder.isCompleted() {
				waitingTime := now - currentOrder.Entry
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = Remove(currentOrders, 0)
				i++
			}
		}
		now++
	}

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
