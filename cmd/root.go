package cmd

import (
	"fmt"
	"weaver/readme"
	"os"
)


func Execute() {
	rm := &readme.Readme{}
	fmt.Printf("%T", rm)
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
	rmFile, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer rmFile.Close()

	//write project name
	// rmFile.WriteString("# " + strings.Title(rm.ProjectName()))
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	readme.WriteProjectName(rm, rmFile)

}
