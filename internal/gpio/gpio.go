package gpio

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func FindDeviceFile() (string, error) {
	// Glob for the sensor directory
	dirs, err := filepath.Glob("/sys/bus/w1/devices/28-*")
	if err != nil {
		return "", err
	}
	if len(dirs) == 0 {
		return "", fmt.Errorf("no DS18B20 devices found")
	}
	// w1_slave file inside the first matching directory
	return filepath.Join(dirs[0], "w1_slave"), nil
}

func readRaw(deviceFile string) ([]string, error) {
	data, err := os.ReadFile(deviceFile)
	if err != nil {
		return nil, err
	}
	// Split into lines
	return strings.Split(string(data), "\n"), nil
}

func ReadTempC(deviceFile string) (float64, error) {
	for {
		lines, err := readRaw(deviceFile)
		if err != nil {
			return 0, err
		}
		// First line ends with "YES" when CRC is good
		if len(lines) > 0 && strings.HasSuffix(lines[0], "YES") {
			// Look for "t=xxxxx" in second line
			parts := strings.Split(lines[1], "t=")
			if len(parts) < 2 {
				return 0, fmt.Errorf("temperature data not found")
			}
			// Parse millidegrees
			raw, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, err
			}
			return float64(raw) / 1000.0, nil
		}
		// CRC failed: wait and retry
		time.Sleep(200 * time.Millisecond)
	}
}
