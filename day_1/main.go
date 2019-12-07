package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic("Failed to open file")
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	fuelNeeded := 0
	for s.Scan() {
		moduleWeight, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		need := moduleWeight/3 - 2
		for {
			if need <= 0 {
				break
			}
			fuelNeeded += need
			need = need/3 - 2
		}

	}
	fmt.Println("Fuel needed:", fuelNeeded)

}
