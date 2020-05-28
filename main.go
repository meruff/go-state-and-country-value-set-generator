package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Country struct {
	Label string
	Value string
}

func main() {
	readFile, err := os.Open("csv/countries.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(readFile).ReadAll()
	if err != nil {
		panic(err)
	}

	allCountries := []*Country{}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := Country{
			Label: line[0],
			Value: line[4],
		}
		allCountries = append(allCountries, &data)
	}

	fileString := `<?xml version="1.0" encoding="UTF-8"?>
<GlobalValueSet xmlns="http://soap.sforce.com/2006/04/metadata">`

	for i, country := range allCountries {
		if i == 0 { // skip first line of .csv (header)
			continue
		}

		fileString += `
	<customValue>
		<fullName>` + after(country.Value, ":") + `</fullName>
		<default>false</default>
		<label>` + country.Label + `</label>
	</customValue>`
	}

	fileString += `
	<masterLabel>Country Picklist</masterLabel>
    <sorted>false</sorted>
</GlobalValueSet>`

	fmt.Println(fileString)
	newFile, err := os.Create("Country_Picklist.globalValueSet")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := newFile.WriteString(fileString)
	if err != nil {
		fmt.Println(err)
		newFile.Close()
		return
	}
	fmt.Println(l)
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
