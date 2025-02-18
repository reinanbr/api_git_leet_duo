package main

import (
	"bytes"

	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sort"
	
)
















func main() {
	user := "example-user"
	startingYear := 2015
	graphs, err := getContributionGraphs(user, startingYear)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully retrieved contribution graphs for", user)
	sortYears := []int{}
	for year := range graphs {
		sortYears = append(sortYears, year)
	}
	sort.Ints(sortYears)
	for _, year := range sortYears {
		fmt.Printf("Year %d: %v\n", year, graphs[year])
	}
}
