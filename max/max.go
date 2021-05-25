//Package max returns the largest of two values.
package max

//Int returns the maximum integer of the two values.
func Int(a, b int) int {
	if a > b {
		return a
	}
	return b
}
