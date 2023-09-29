// converter/converter.go
package converter

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func ConvertYAMLToJSON(inputFile, outputFile string) error {
	// Read YAML input
	yamlData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %s", err)
	}

	// Unmarshal YAML using map[interface{}]interface{}
	var data map[interface{}]interface{}
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML: %s", err)
	}

	// Convert map[interface{}]interface{} to a more JSON-friendly structure
	jsonData, err := convertToJSONFriendlyStructure(data)
	if err != nil {
		return fmt.Errorf("error converting to JSON-friendly structure: %s", err)
	}

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %s", err)
	}

	// Write JSON to the output file
	err = os.WriteFile(outputFile, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %s", err)
	}

	return nil
}

func convertToJSONFriendlyStructure(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[interface{}]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			keyString, ok := key.(string)
			if !ok {
				return nil, fmt.Errorf("non-string key found in YAML: %v", key)
			}
			convertedValue, err := convertToJSONFriendlyStructure(value)
			if err != nil {
				return nil, err
			}
			result[keyString] = convertedValue
		}
		return result, nil
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, item := range v {
			convertedItem, err := convertToJSONFriendlyStructure(item)
			if err != nil {
				return nil, err
			}
			result[i] = convertedItem
		}
		return result, nil
	default:
		return v, nil
	}
}
