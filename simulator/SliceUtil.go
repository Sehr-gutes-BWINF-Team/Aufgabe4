package simulator

import "errors"

// CalculateAverage of a slice full of integers.
func CalculateAverage(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	average := total / len(numbers) // len  function return array size
	return average
}

// RemoveByIndex removes an order at a specific index.
func RemoveByIndex(s int, slice []Order) []Order {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveOrder(order Order, orders []Order) []Order {
	index, _ := GetIndexByOrderId(order.Id, orders)
	return RemoveByIndex(index, orders)
}

func GetIndexByOrderId(orderId int, orders []Order) (int, error) {
	for i := range orders {
		if orders[i].Id == orderId {
			return i, nil
		}
	}
	return -1, errors.New("not found")
}

func GetOrderById(orderId int, orders []Order) (Order, error) {
	index, err := GetIndexByOrderId(orderId, orders)
	return orders[index], err
}

func LowestTimeOrder(orders []Order) Order {
	var minOrder = orders[0]
	for _, value := range orders {
		if minOrder.Time > value.Time {
			minOrder = value
		}
	}
	return minOrder
}
