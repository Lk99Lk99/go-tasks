package listprops

// ContainsValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert true, falls v in der Liste enthalten ist, sonst false.
func ContainsValue(list []int, v int) bool {
	// TODO

	for i := 0; i < len(list); i++ {
		if v == list[i] {
			return true
		}
	}
	return false
}
