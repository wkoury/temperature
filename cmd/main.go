package main

import (
	"fmt"
	"log"
	"temperature/internal/gpio"
	"time"
)

func main() {
	deviceFile, err := gpio.FindDeviceFile()
	if err != nil {
		log.Fatalf("Error finding device: %v", err)
	}

	fmt.Printf("Reading temperature from %s\n", deviceFile)
	for {
		tempC, err := gpio.ReadTempC(deviceFile)
		if err != nil {
			log.Printf("Error reading temperature: %v", err)
		} else {
			fmt.Printf("Temperature: %.3f °C\n", tempC)
		}
		time.Sleep(1 * time.Second)
	}
}
