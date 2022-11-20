package simulator

import (
	"errors"
)

// CalculateAverage of a slice full of integers.
func CalculateAverage(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	average := total / len(numbers) // len function return array size
	return average
}

// RemoveByIndex removes an order at a specific index.
func RemoveByIndex(s int, slice []Order) []Order {
	return append(slice[:s], slice[s+1:]...)
}

// GetHighestInt of a slice full of integers.
func GetHighestInt(numbers []int) int {
	highest := -1
	for _, number := range numbers {
		if highest == -1 {
			highest = number
		}
		if highest < number {
			highest = number
		}
	}
	return highest
}

// GetLowestInt of a slice full of integers.
func GetLowestInt(numbers []int) int {
	var highest = -1
	for _, number := range numbers {
		if highest == -1 {
			highest = number
		}
		if highest > number {
			highest = number
		}
	}
	return highest
}

func RemoveOrder(order Order, orders []Order) []Order {
	index, _ := GetIndexByOrderID(order.ID, orders)
	return RemoveByIndex(index, orders)
}

func GetIndexByOrderID(orderId int, orders []Order) (int, error) {
	for i := range orders {
		if orders[i].ID == orderId {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func LowestTimeOrder(orders []Order) Order {
	var minOrder = orders[0]
	for _, value := range orders {
		if value.RequiredTime < minOrder.RequiredTime {
			minOrder = value
		}
	}
	return minOrder
}
