/*
Copyright © 2024 David Adediji <idavid.adediji@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the picture based on the id passed",
	Long:  `Saves a picture to your directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "0"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}
		URL := "https://picsum.photos/id/" + gopherName + "/5000/3333"
		fmt.Println("Try to get '" + gopherName + "' Gopher...")

		// Retrieve the gopher images
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(gopherName + ".jpg")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			// Copy contents of response to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("File saved in '" + out.Name() + "'!")
		} else {
			fmt.Println("Error: " + gopherName + " not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
