package simulator

import (
	"log"
)

func FirstSimulation(orders []Order) {
	Simulate(len(orders), orders, func(currentOrder *Order, currentOrders []Order) *Order {
		currentOrder = &currentOrders[0]
		currentOrder.WorkOn()
		return currentOrder
	})
}

func SecondSimulation(orders []Order) {
	Simulate(len(orders), orders, func(currentOrder *Order, currentOrders []Order) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(LowestTimeOrder(currentOrders).ID, currentOrders)
		currentOrder = &currentOrders[currentOrderIndex]
		currentOrder.WorkOn()
		return currentOrder
	})
}

// Simulate simulates the bike repair shop scenario.
func Simulate(simulationLength int, orders []Order, simulation func(currentOrder *Order, currentOrders []Order) *Order) {
	log.Println("Started simulation!")

	var currentOrders []Order
	var finishedOrders []int
	completedOrders, overallTime, workTime := 0, 0, 0
	for completedOrders < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 && workTime <= 7 {
			currentOrder = simulation(currentOrder, currentOrders)

			if currentOrder.IsCompleted() {
				waitingTime := overallTime - currentOrder.Entry
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = RemoveOrder(*currentOrder, currentOrders)
				completedOrders++
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
