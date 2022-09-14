package main

import (
	"bufio"
	"bwinf22/simulator"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("resources/fahrradwerkstatt0.txt")

	if err != nil {
		fmt.Println(err)
	}

	var orders []simulator.Order
	bScanner := bufio.NewScanner(readFile)
	bScanner.Split(bufio.ScanLines)

	for bScanner.Scan() {
		line := bScanner.Text()
		orderData := strings.Split(line, " ")

		if len(orderData) != 2 {
			break
		}

		entry, _ := strconv.Atoi(orderData[0])
		time, _ := strconv.Atoi(orderData[1])
		orders = append(orders, simulator.Order{Entry: entry, Time: time, Completion: time})
	}

	simulator.Simulate(len(orders), orders)

	err = readFile.Close()
	if err != nil {
		return
	}
}
