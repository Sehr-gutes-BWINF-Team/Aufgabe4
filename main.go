package main

import (
	"bufio"
	"bwinf22/simulator"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var orders = getOrdersFromFile("resources/fahrradwerkstatt4.txt")
	simulator.FirstInFirstOut(orders)
	simulator.CompleteThenShortest(orders)
	simulator.RoundRobin(orders)
}

func getOrdersFromFile(filePath string) []simulator.Order {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	var orders []simulator.Order
	bScanner := bufio.NewScanner(readFile)
	bScanner.Split(bufio.ScanLines)
	var id int

	for bScanner.Scan() {
		line := bScanner.Text()
		orderData := strings.Split(line, " ")

		if len(orderData) != 2 {
			break
		}

		entry, _ := strconv.Atoi(orderData[0])
		time, _ := strconv.Atoi(orderData[1])
		orders = append(orders, simulator.Order{ID: id, EntryTime: entry, RequiredTime: time, TimeLeftUntilCompletion: time})
		id++
	}

	err = readFile.Close()
	if err != nil {
		return nil
	}

	return orders
}
