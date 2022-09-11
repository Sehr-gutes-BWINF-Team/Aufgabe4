package main

import (
	"bufio"
	"bwinf22/lib"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type order struct {
	entry int
	time  int
}

func main() {
	readFile, err := os.Open("resources/fahrradwerkstatt0.txt")

	if err != nil {
		fmt.Println(err)
	}

	var orders []order
	bScanner := bufio.NewScanner(readFile)
	bScanner.Split(bufio.ScanLines)
	line := ""

	for bScanner.Scan() {
		line = bScanner.Text()

		orderData := strings.Split(line, " ")

		if len(orderData) != 2 {
			break
		}

		entry, _ := strconv.Atoi(orderData[0])
		time, _ := strconv.Atoi(orderData[1])
		orders = append(orders, order{entry, time})
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].time < orders[j].time
	})

	lib.Simulate(len(orders))

	err = readFile.Close()
	if err != nil {
		return
	}
}
