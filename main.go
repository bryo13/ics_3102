package main

import (
	"fmt"
	"os"
	"strings"
)

type propertyCount struct {
	uri      string
	count    int
	mostUsed bool
}

func main() {
	// unzip the file
	// Read the Turtle file assuming it is in the same directory as this script
	data, err := os.ReadFile("mappingbased_objects_en.ttl")
	if err != nil {
		panic(err)
	}

	// Initialize variables
	properties := map[string]*propertyCount{}
	mostUsedProperty := &propertyCount{}

	// Extract property data
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Fields(line)

		if len(parts) < 3 || parts[0][0] == '#' || parts[0] == "" {
			continue
		}

		predicate := strings.TrimSuffix(parts[1], ":")

		count, ok := properties[predicate]
		if !ok {
			count = &propertyCount{uri: predicate}
			properties[predicate] = count
		}
		count.count++

		if count.count > mostUsedProperty.count {
			mostUsedProperty.uri = count.uri
			mostUsedProperty.count = count.count
			mostUsedProperty.mostUsed = true
		}
	}

	// Print results
	fmt.Printf("Number of unique properties: %d\n", len(properties))
	fmt.Printf("Most frequently used property: %s (%d occurrences)\n", mostUsedProperty.uri, mostUsedProperty.count)
}
