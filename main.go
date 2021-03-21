package main

import (
	"fmt"

	"github.com/gain620/go-portmap/port"
)

func main() {
	fmt.Println("Scanning for ports...")

	// open := port.ScanPort("tcp", "localhost", 80)
	// fmt.Printf("Port open : %t\n", open)

	results := port.WellKnownScan("tcp", "localhost")
	for _, v := range results {
		fmt.Println("PORT:", v.Port, " | STATE:", v.State)
	}
}
