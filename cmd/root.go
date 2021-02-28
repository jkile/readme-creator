package cmd

import (
	"fmt"
	"weaver/readme"
)


func Execute() {
	rm := readme.Readme{}
	rm.SetProjectName(prompt("Enter project name: "))
	rm.SetImageUrl(prompt("Enter project image path: "))
	rm.SetDescription(prompt("Enter project description: "))
	rm.SetUsage(prompt("Enter project usage: "))
	rm.SetContributing(prompt("Enter how to contribute: "))
	fmt.Printf("\nProject Name: %v\nImage Path: %v\nDescription: %v\nUsage: %v\nContributing: %v",
		rm.ProjectName(),
		rm.ImageUrl(),
		rm.Description(),
		rm.Usage(),
		rm.Contributing(),
	)
}
