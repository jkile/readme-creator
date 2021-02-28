package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func prompt(prompt string) string{
	fmt.Print(prompt)
	var response string
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if strings.Contains(text, "\n") {
			text = strings.Replace(text, "\n", "", -1)
			//fmt.Printf("You entered: %v%v%v\n", colors.Blue, text, colors.Reset)
			response = text
			break
		}
	}
	return response
}