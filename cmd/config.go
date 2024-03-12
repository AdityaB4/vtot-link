/*
Copyright © 2024 Aditya Bajaj <bajaj.aditya2204@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigOption int

// TODO: Find a better way for this (with intellisense)
const (
	API_KEY ConfigOption = iota
	THRESHOLD
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config [--api-key a] [--threshold t]",
	Aliases: []string{"c"},
	Short:   "used to configure the VirusTotal api key",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println("No config options given. Maybe this will help you:")
			fmt.Println(cmd.Usage())
			return
		}
		if apiKey, err := cmd.Flags().GetString("api-key"); err == nil && apiKey != "" {
			setConfigOption(API_KEY, apiKey)
			fmt.Println("api key set to", apiKey)
		}
		// TODO: Check negative
		if threshold, err := cmd.Flags().GetInt("threshold"); err == nil {
			setConfigOption(THRESHOLD, threshold)
			fmt.Printf("threshold set to %d seconds\n", threshold)
		}
	},
}

func setConfigOption(option ConfigOption, value any) {
	switch option {
	case API_KEY:
		viper.Set("apiKey", value)
	case THRESHOLD:
		viper.Set("threshold", value)
	default:
		fmt.Println("Invalid config option provided. Aborting")
	}
	writeConfigFile()
}

func writeConfigFile() {
	if err := viper.WriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// If the config file doesn't exist, try to create it
			viper.SafeWriteConfig()
		} else {
			// Handle other errors
			fmt.Println("Error writing config:", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().String("api-key", "", "--api-key <your VT api key>")
	configCmd.Flags().String("threshold", "", "--threshold <desired threshold in seconds>")
}
