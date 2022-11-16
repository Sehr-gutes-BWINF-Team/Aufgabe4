package simulator

import (
	"log"
)

func FirstSimulation(orders []Order) {
	Simulate(len(orders), orders, func(currentOrder *Order, currentOrders []Order, overallTime int) *Order {
		currentOrder = &currentOrders[0]
		return currentOrder
	})
}

func SecondSimulation(orders []Order) {
	Simulate(len(orders), orders, func(currentOrder *Order, currentOrders []Order, overallTime int) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(LowestTimeOrder(currentOrders).ID, currentOrders)
		currentOrder = &currentOrders[currentOrderIndex]
		return currentOrder
	})
}

func ThirdSimulation(orders []Order) {
	Simulate(len(orders), orders, func(currentOrder *Order, currentOrders []Order, overallTime int) *Order {
		currentOrderIndex, _ := GetIndexByOrderID(HighestPriorityOrder(currentOrders, overallTime).ID, currentOrders)
		currentOrder = &currentOrders[currentOrderIndex]
		return currentOrder
	})
}

// Simulate simulates the bike repair shop scenario.
func Simulate(simulationLength int, orders []Order, simulation func(currentOrder *Order, currentOrders []Order, overallTime int) *Order) {
	log.Println("Started simulation!")

	var currentOrders []Order
	var finishedOrders []int
	completedOrders, overallTime, dayTime := 0, 0, 0
	for completedOrders < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 && dayTime <= 17*60 && dayTime >= 9*60 {
			currentOrder = simulation(currentOrder, currentOrders, overallTime)
			currentOrder.WorkOn()

			if currentOrder.IsCompleted() {
				waitingTime := overallTime - currentOrder.Entry
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = RemoveOrder(*currentOrder, currentOrders)
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
	log.Println("Highest waiting Time: ", highest, "minutes.")
	log.Println("Lowest waiting Time: ", lowest, "minutes.")
	log.Println("Average:", average, "minutes.")
	log.Println("Which are equal to", average/60, "hours.")
	log.Println("Or", average/60/24, "days.")
	log.Println("Stopped simulation!")
}

var x = 0

func Test(simulationLength int, orders []Order) {
	log.Println("Started simulation!")

	var currentOrders []Order
	var finishedOrders []int
	completedOrders, overallTime, workTime := 0, 0, 0
	for completedOrders < simulationLength {
		currentOrders = append(currentOrders, getNewOrders(overallTime, orders)...)
		var currentOrder *Order

		if len(currentOrders) != 0 && workTime <= 7*60 {

			if x == len(currentOrders) || x > len(currentOrders) {
				x = 0
			}

			currentOrder = &currentOrders[x]
			currentOrder.WorkOn()
			x++

			if currentOrder.IsCompleted() {
				waitingTime := overallTime - currentOrder.Entry
				finishedOrders = append(finishedOrders, waitingTime)
				currentOrders = RemoveOrder(*currentOrder, currentOrders)
				completedOrders++
			}
		}

		workTime++

		if workTime == 24*60 {
			workTime = 0
		}

		overallTime++
	}

	average := CalculateAverage(finishedOrders)
	lowest := GetLowestInt(finishedOrders)
	highest := GetHighestInt(finishedOrders)
	log.Println("Highest waiting Time: ", highest, "minutes.")
	log.Println("Lowest waiting Time: ", lowest, "minutes.")
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
