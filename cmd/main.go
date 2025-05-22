package main

import (
	"context"
	"fmt"
	"log"
	"temperature/internal/db"
	"temperature/internal/gpio"
	"temperature/internal/temp"
	"time"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	db.Init()

	deviceFile, err := gpio.FindDeviceFile()
	if err != nil {
		log.Fatalf("Error finding device: %v", err)
	}

	ctx := context.Background()

	fmt.Printf("Reading temperature from %s\n", deviceFile)
	for {
		tempC, err := gpio.ReadTempC(deviceFile)
		if err != nil {
			log.Printf("Error reading temperature: %v", err)
			time.Sleep(10 * time.Minute)
			continue
		}

		fmt.Printf("Temperature: %.3f °F\n", temp.CtoF(tempC))
		err = db.InsertTemperatureRow(ctx, tempC, time.Now())
		if err != nil {
			log.Printf("Error inserting temperature row: %v", err)
			time.Sleep(10 * time.Minute)
			continue
		}

		time.Sleep(5 * time.Minute)
	}
}
