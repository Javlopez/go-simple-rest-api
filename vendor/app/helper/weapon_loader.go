package helper

import (
	"os"
	"app/model"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
	"fmt"
)

func LoadWeaponsCsv() []model.Weapon {

	var weapons []model.Weapon

	//Name,Rarity,Fire Rate,Mag Size,Reload Time,Multiplier,Damage (Body),Damage (Head),Fall Off,DPS (Body),DPS (Head),
	// Shots/Kill Avg (Body),Shots/Kill Avg (Head)
	csvFile, _ := os.Open("fortnite-weapons.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		fireRate, err := strconv.ParseFloat(line[2], 64)
		magSize, err := strconv.Atoi(line[3])
		reloadTime, err := strconv.ParseFloat(line[4], 64)
		multiplier, err := strconv.Atoi(line[5])
		damageBody, err := strconv.Atoi(line[6])
		damageHead, err := strconv.Atoi(line[7])
		dpsBody, err := strconv.Atoi(line[9])
		dpsHead, err := strconv.Atoi(line[10])
		shotsKillAvgBody, err := strconv.ParseFloat(line[11], 32)
		shotsKillAvgHead, err := strconv.ParseFloat(line[12], 32)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		weapons = append(weapons, model.Weapon{
			Name: 		line[0],
			Rarity: 	line[1],
			FireRate:	fireRate,
			MagSize: 	magSize,
			ReloadTime:       reloadTime,
			Multiplier:       multiplier,
			DamageBody:       damageBody,
			DamageHead:       damageHead,
			FallOff:          line[8],
			DPSBody:          dpsBody,
			DPSHead:          dpsHead,
			ShotsKillAvgBody: shotsKillAvgBody,
			ShotsKillAvgHead: shotsKillAvgHead,
		})
	}
	return weapons
}
