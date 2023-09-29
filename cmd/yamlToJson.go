/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/vickyonit/toto/pkg/converter"

	"github.com/spf13/cobra"
)

// jsonToYamlCmd represents the jsonToYaml command
var yamlToJsonCmd = &cobra.Command{
	Use:   "yaml-to-json",
	Short: "Convert YAML to JSON",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile := args[1]
		// Perform YAML to JSON conversion
		err := converter.ConvertYAMLToJSON(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Conversion successful. JSON written to %s\n", outputFile)
	},
}

func init() {
	convertCmd.AddCommand(yamlToJsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yamlToJsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yamlToJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
