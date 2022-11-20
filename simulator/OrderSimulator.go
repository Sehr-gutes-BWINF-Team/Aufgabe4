package simulator

import (
	"log"
)

func FirstInFirstOut(orders []Order) {
	Simulate(orders, func(currentOrder *Order, currentOrders []Order, latestIndex *int, time int) *Order {
		return currentOrder
	}, func(currentOrders []Order) *Order {
		return &currentOrders[0]
	})
}

func RoundRobin(orders []Order) {
	Simulate(orders, func(currentOrder *Order, currentOrders []Order, latestIndex *int, time int) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(LowestTimeOrder(currentOrders).ID, currentOrders)
		return &currentOrders[currentOrderIndex]
	}, func(currentOrders []Order) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(LowestTimeOrder(currentOrders).ID, currentOrders)
		return &currentOrders[currentOrderIndex]
	})
}

func CompleteThenShortest(orders []Order) {
	Simulate(orders, func(currentOrder *Order, currentOrders []Order, latestIndex *int, time int) *Order {
		return currentOrder
	}, func(currentOrders []Order) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(LowestTimeOrder(currentOrders).ID, currentOrders)
		return &currentOrders[currentOrderIndex]
	})
}

// Simulate simulates the bike repair shop scenario.
func Simulate(
	orders []Order,
	changeCurrentOrder func(currentOrder *Order, currentOrders []Order, latestIndex *int, time int) *Order,
	getNextOrder func(currentOrders []Order) *Order,
) {
	log.Println("Started Simulation!")

	var currentOrders []Order
	var finishedOrders []int
	counter := 0
	completedOrders, overallTime, dayTime := 0, 0, 0
	var currentOrderPointer *Order = nil
	for completedOrders < len(orders) {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		if len(currentOrders) != 0 && dayTime <= 17*60 && dayTime >= 9*60 {
			if currentOrderPointer == nil {
				currentOrderPointer = getNextOrder(currentOrders)
			} else {
				currentOrderPointer = changeCurrentOrder(currentOrderPointer, currentOrders, &counter, overallTime)
			}
			currentOrderPointer.WorkOn()

			if currentOrderPointer.IsCompleted() {
				waitingTime := overallTime - currentOrderPointer.EntryTime
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = RemoveOrder(*currentOrderPointer, currentOrders)
				currentOrderPointer = nil
				completedOrders++
			}
		}

		dayTime++

		if dayTime == 24*60 {
			dayTime = 0
		}

		overallTime++
	}
	// log.Println(overallTime)
	// log.Println(finishedOrders)
	average := CalculateAverage(finishedOrders)
	lowest := GetLowestInt(finishedOrders)
	highest := GetHighestInt(finishedOrders)
	log.Println("Highest waiting RequiredTime:", highest, "minutes.")
	log.Println("Lowest waiting RequiredTime:", lowest, "minutes.")
	log.Println("Average waiting time:", average, "minutes =", average/60, "hours =", average/60/24, "days")

	log.Println("Stopped Simulation!")
}

func getNewOrders(now int, orders []Order) []Order {
	var newOrders []Order
	for i := range orders {
		if now == orders[i].EntryTime {
			newOrders = append(newOrders, orders[i])
		}
	}
	return newOrders
}
