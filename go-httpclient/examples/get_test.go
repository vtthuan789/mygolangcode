package examples

import (
	"fmt"
	"testing"
)

func Test_GetEndpoints(t *testing.T) {
	// Initializtion

	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	fmt.Println("Error:", err)
	fmt.Println("Endpoints:", endpoints)

}
