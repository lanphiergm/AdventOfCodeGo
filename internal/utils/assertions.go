// Contains utility functions useful across puzzles and years
package utils

import "testing"

// Asserts that the two values are equal and if not, fails the test with a
// useful message.
func AssertAreEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Fatalf("Expected: <%v>, Actual: <%v>", expected, actual)
	}
}
