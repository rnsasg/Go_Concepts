package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func validateAndParseConfiguration(configStr string) []string {
	// Split the configurations
	configurations := strings.Split(configStr, "|")

	// Regular expression to match valid configuration strings
	regex := regexp.MustCompile(`^(\d{4})([a-zA-Z0-9]{10})$`)

	// Maps to track indices and configuration values
	indicesMap := make(map[string]bool)
	valuesMap := make(map[string]bool)

	// Slice to store the configurations
	configPairs := make([][2]string, len(configurations))

	for i, config := range configurations {
		// Validate each configuration string using regex
		matches := regex.FindStringSubmatch(config)
		if matches == nil {
			return []string{"Invalid configuration"}
		}

		ordinalIndex := matches[1]
		configValue := matches[2]

		// Check for "0000" ordinal index
		if ordinalIndex == "0000" {
			return []string{"Invalid configuration"}
		}

		// Check for unique ordinal indices
		if indicesMap[ordinalIndex] {
			return []string{"Invalid configuration"}
		}
		indicesMap[ordinalIndex] = true

		// Check for unique configuration values
		if valuesMap[configValue] {
			return []string{"Invalid configuration"}
		}
		valuesMap[configValue] = true

		// Store the valid configuration pair
		configPairs[i] = [2]string{ordinalIndex, configValue}
	}

	// Sort the configurations based on the ordinal index
	sort.Slice(configPairs, func(i, j int) bool {
		return configPairs[i][0] < configPairs[j][0]
	})

	// Check for missing indices
	for i := 0; i < len(configPairs); i++ {
		expectedIndex := fmt.Sprintf("%04d", i+1)
		if configPairs[i][0] != expectedIndex {
			return []string{"Invalid configuration"}
		}
	}

	// Extract the configuration values in the correct order
	result := make([]string, len(configPairs))
	for i, pair := range configPairs {
		result[i] = pair[1]
	}

	return result
}

func main() {
	// Test examples
	configStr1 := "0001f7c22e7904|000276a3a4d214|000365d29f4a4b"
	configStr2 := "0002f7c22e7904|000176a3a4d214|000365d29f4a4b"

	fmt.Println(validateAndParseConfiguration(configStr1)) // ["f7c22e7904", "76a3a4d214", "65d29f4a4b"]
	fmt.Println(validateAndParseConfiguration(configStr2)) // ["Invalid configuration"]
}
