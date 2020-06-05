package main

import (
	"fmt"
)

func main() {
	c := saveData(fetchData(
		prepareData(
			generateData(),
		),
	))

	for data := range c {
		fmt.Printf("Items Saved: %+v \n", data)
	}
}
