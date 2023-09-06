// Package Declaration
package mascot_test

// Import required packages
import (
	"testing"

	"example.com/p01/mascot"
)

// Function: TestMascot()
// Purpose: To test that the BestMascot() function returns
// Returns: Error message to console
func TestMascot(t *testing.T) {
	// If fucntion to test if the function outputs the correct mascot
	if mascot.BestMascot() != "Maverick T. Mustang" {
		t.Fatal("Wrong Mascot :0")
	}
}
