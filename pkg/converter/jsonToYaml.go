package to_yaml

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ConvertJSONToYAML(inputFile, outputFile string) error {
	// Read JSON input
	jsonData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %s", err)
	}

	// Unmarshal JSON
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %s", err)
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(jsonMap)
	if err != nil {
		return fmt.Errorf("error marshalling to YAML: %s", err)
	}

	// Write YAML to the output file
	err = os.WriteFile(outputFile, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing YAML to file: %s", err)
	}

	return nil
}
