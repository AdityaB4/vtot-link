package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "./config.yaml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vtot-link",
	Short: "A CLI tool to query VirusTotal's link API with 1 command",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	// TODO: Think about whether a flag is needed for this
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "config file")

	// Flags scoped only to this command
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
