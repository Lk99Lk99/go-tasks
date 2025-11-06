package listprops

// CountValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert die Anzahl der Vorkommen von v in der Liste.
func CountValue(list []int, v int) int {
	// TODO
	a := 0
	for i := 0; i < len(list); i++ {
		if v == list[i] {
			a = a + 1
		}
	}

	return a
}
