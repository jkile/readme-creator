package readme

import (
	"strings"
	"fmt"
	"os"
)

type File = os.File

func WriteProjectName (data *Readme, file *File) bool {
	title := "# " + strings.Title(data.ProjectName())
	_, err := file.WriteString(title)
	if err != nil {
		fmt.Println("Error Writing Project Name: ", err)
		return false
	}
	return true
}