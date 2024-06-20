package main

import (
	"fmt"
)

func classifyTriangleAngle(angle1, angle2 int) string {
	// Calculate the missing angle
	missingAngle := 180 - (angle1 + angle2)

	// Determine the type of the missing angle
	if missingAngle < 90 {
		return "Acute"
	} else if missingAngle == 90 {
		return "Right"
	} else {
		return "Obtuse"
	}
}

func main() {
	var angle1, angle2 int
	fmt.Scan(&angle1, &angle2)
	fmt.Println(classifyTriangleAngle(angle1, angle2))
}
