package main

import "fmt"

func main() {

	// Fragment 1
	wasser := 0
	limit := 10

	fmt.Println("--- Start: Bewässerung ---")
	for wasser < limit {
		fmt.Println("Gieße magisches Wasser...")
		fmt.Println("Limit: ", limit)
		wasser++ // FIX
	}
	fmt.Println("Bewässerung abgeschlossen! \n")

	// Fragment 2
	lichtEnergie := 100

	fmt.Println("--- Start: Licht-Zufuhr ---")
	for lichtEnergie > 0 {
		fmt.Println("Lichtstrahl wird fokussiert...")
		lichtEnergie-- // FIX
		fmt.Println("Lichtenergie: ", lichtEnergie)
	}
	fmt.Println("Lichtenergie verbraucht! \n")

	// Fragment 3
	istGezähmt := false
	versuche := 0

	fmt.Println("--- Start: Wächter-Check ---")
	for !istGezähmt {
		fmt.Printf("Versuch %d: Besänftige die Ranken...\n", versuche)
		fmt.Println("Versuche: ", versuche)
		versuche++

		if versuche >= 10 { // FIX
			istGezähmt = true
		}
	}

	fmt.Println("Mission abgeschlossen: Die Ranken sind gezähmt!")
}
