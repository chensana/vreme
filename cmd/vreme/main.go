package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chensana/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Upravljanje vremenskih postaj in meritev",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve (°C)",
				Value: -50.0,
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve (°C)",
				Value: 50.0,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			vreme.DodajPostajo("KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo("MB01", "Maribor Tabor", "Maribor")

			vreme.DodajMeritev("MB01", 3, minTemp, maxTemp)
			vreme.DodajMeritev("MB01", 60, minTemp, maxTemp) // izven mej
			vreme.DodajMeritev("XYZ", 20, minTemp, maxTemp) // neobstoječa postaja

			fmt.Println()
			vreme.IzpisMeritev()
			fmt.Println()
			vreme.IzpisPorocila(stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}