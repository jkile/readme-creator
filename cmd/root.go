package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use: "weaver",
	Short: "Weaver makes your readme like a basket",
	Long: `A fast readme file generator to speed up the process of 
			documenting projects and having to manually markup files`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Please enter something: ")
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			if strings.Contains(text, "\n") {
				text = strings.Replace(text, "\n", "", -1)
				fmt.Println("You entered: ", text)
			}
		}
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil{
		log.Fatalf("Fatal error: %s", err)
	}
}

//func init() {
//	cobra.on
//}