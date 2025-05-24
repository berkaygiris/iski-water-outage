package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/berkaygiris/iski-water-outage/internal"
)

func main() {
	apiResponse, err := internal.FetchData()
	if err != nil {
		log.Fatalf("Error fetching the data: %v", err)
	}

	outages, err := internal.ParseOutages(apiResponse)
	if err != nil {
		log.Fatalf("Error parsing outages: %v", err)
	}

	for _, o := range outages {
		fmt.Printf("ID: %s\n", o.ID)
		fmt.Printf("Start Date: %s\n", o.StartDate)
		fmt.Printf("End Date: %s\n", o.EndDate)
		fmt.Printf("Info: %s\n", o.Info)
		fmt.Println("Zones:")
		for _, z := range o.Zones {
			fmt.Printf("  - %s - %s\n", z.District, strings.Join(z.Neighborhoods, ", "))
		}
		fmt.Println()
	}
}
