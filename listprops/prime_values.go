package listprops

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	// TODO

	v := 0

	for i := 1; i < len(list)+1; i++ {
		if IsPrime(i) {
			v++
		}
	}
	return v

}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
