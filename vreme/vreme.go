// Paket vreme ponuja podatkovno strukturo za vremenske postaje
// ter funkcije za dodajanje postaj, vnos meritev in izpis poročil.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z osnovnimi podatki in meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje = make(map[string]Postaja)

// DodajPostajo v slovar postaj doda novo postajo z dano oznako, imenom in krajem.
func DodajPostajo(oznakaP, imeP, krajP string) {
	novaPostaja := Postaja{ime: imeP, kraj: krajP}
	postaje[oznakaP] = novaPostaja
}

// DodajMeritev doda temperaturo k postaji z dano oznako, če je znotraj mej [minTemp, maxTemp].
func DodajMeritev(oznaka string, temperatura, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Temperatura ni v dovoljenem obmocju.")
		return
	}

	p, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Postaja ni veljavna.")
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
}

func statistika(oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	v, exists := postaje[oznaka]
	if !exists || len(v.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = v.meritve[0]
	najvisja = v.meritve[0]
	povprecna = 0.0

	for _, stevilo := range v.meritve {
		if stevilo < najnizja {
			najnizja = stevilo
		}
		if stevilo > najvisja {
			najvisja = stevilo
		}
		povprecna += stevilo
	}
	povprecna = povprecna / float64(len(v.meritve))
	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše vse zbrane meritve za vse postaje.
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for i, v := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", i, v.ime, v.kraj, v.meritve)
	}
}

// IzpisPorocila izpiše statistično poročilo in temperaturni opis za postaje z vsaj stMeritev meritvami.
func IzpisPorocila(stMeritev int) {
	for i, v := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(i, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", v.ime)
		} else {
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprecje %.1f °C, -> ", v.ime, najnizja, najvisja, povprecna)
			switch {
			case povprecna >= 25.0:
				fmt.Printf("Vroče\n")
			case povprecna >= 15.0:
				fmt.Printf("Toplo\n")
			case povprecna >= 5.0:
				fmt.Printf("Hladno\n")
			default:
				fmt.Printf("Mrzlo\n")
			}
		}
	}
}