package main

import (
	"encoding/json"
	"fmt"
)

type Deployment struct {
	DeploymentID string `json:"deployment_id"`
	Status       string `json:"status"`
}

func isValidDeploymentID(deploymentID string) bool {
	if len(deploymentID) != 12 || deploymentID[:2] != "d-" {
		return false
	}
	for _, ch := range deploymentID[2:] {
		if !('a' <= ch && ch <= 'z') && !('0' <= ch && ch <= '9') {
			return false
		}
	}
	return true
}

func evaluate(d []string) []int32 {
	successCount := int32(0)
	failCount := int32(0)
	errorCount := int32(0)

	for _, item := range d {
		var deployment Deployment
		err := json.Unmarshal([]byte(item), &deployment)
		if err != nil || !isValidDeploymentID(deployment.DeploymentID) {
			errorCount++
			continue
		}

		switch deployment.Status {
		case "Success":
			successCount++
		case "Fail":
			failCount++
		default:
			errorCount++
		}
	}

	return []int32{successCount, failCount, errorCount}
}

func main() {
	deployments := []string{
		`{"deployment_id": "d-123456abcd", "status": "Success"}`,
		`{"deployment_id": "d-098765efgh", "status": "Fail"}`,
		`{"deployment_id": "d-invalid123", "status": "Success"}`,
		`{"deployment_id": "d-12345efghi", "status": "Unknown"}`,
		`{"deployment_id": "d-123456abc", "status": "Fail"}`,
	}

	results := evaluate(deployments)
	fmt.Println(results) // Output: [1, 1, 3]
}
