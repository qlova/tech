//Package min returns the smallest of two values.
package min

//Int returns the minumum integer of the two values.
func Int(a, b int) int {
	if a < b {
		return a
	}
	return b
}
