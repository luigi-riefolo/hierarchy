package main

import (
	"fmt"
	"log"

	"github.com/luigi-riefolo/hierarchy"
)

func main() {

	dataFile := "data/complete.json"
	hierarchy, err := hierarchy.NewHierarchy(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	employeeOne, employeeTwo := "Q", "T1"
	fmt.Printf("Finding manager for \"%s\" and \"%s\"\n", employeeOne, employeeTwo)
	manager, err := hierarchy.FindClosestManager(employeeOne, employeeTwo)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Found manager: ", manager.ID)
	}
}
