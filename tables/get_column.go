package tables

// GetColumn erwartet ein zweidimensionales Array (Tabelle) und eine Spaltennummer.
// Liefert eine Liste mit den Werten der angegebenen Spalte.
// Falls die Zeilen unterschiedliche Längen haben, wird für fehlende Werte ein leerer String geliefert.
func GetColumn(table [][]string, col int) []string {
	// TODO

	var result []string

	for _, row := range table {
		if col < len(row) {
			result = append(result, row[col])
		} else {
			result = append(result, "")
		}
	}

	return result
}
