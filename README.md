##Vremenska postaja##

Vreme
Modul v programskem jeziku Go za upravljanje vremenskih postaj in temperaturnih meritev.

Opis
Projekt omogoča:

dodajanje vremenskih postaj,
dodajanje temperaturnih meritev,
izpis vseh meritev,
izračun statistike (minimum, maksimum, povprečje),
izpis poročila o vremenskih postajah.
Projekt je izdelan v okviru domače naloge pri predmetu Programski jezik Go.

Struktura projekta
.
├── cmd/
│   └── vreme/
│       └── main.go
├── vreme/
│   └── vreme.go
├── go.mod
└── go.sum
Namestitev
Prenos odvisnosti:

go mod tidy
Zagon
go run ./cmd/vreme
Primer zagona z argumenti:

go run ./cmd/vreme --minTemp -20 --maxTemp 40 --stMeritev 2
Stikala ukazne vrstice
Stikalo	Opis	Privzeta vrednost
--minTemp	Najnižja dovoljena temperatura	-50
--maxTemp	Najvišja dovoljena temperatura	50
--stMeritev	Najmanjše število meritev za statistiko	3
Paket vreme
Paket vreme izvaža naslednje funkcije:

DodajPostajo
DodajMeritev
IzpisMeritev
IzpisPorocila
Funkcija statistika je namenjena interni uporabi in ni izvožena.

Gradnja
go build -o vreme_app ./cmd/vreme

Verzija
v1.0.0
