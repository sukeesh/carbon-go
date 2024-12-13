package main

import (
	"errors"
	"fmt"
)

// Defaults for estimation
const (
	defaultCPUPowerKW      = 0.05  // CPU power usage in kilowatts (50W)
	defaultCarbonIntensity = 475.0 // Global average carbon intensity in gCO2/kWh
)

// Config holds configurable parameters for CO2 estimation
type Config struct {
	CPUPowerKW      float64 // CPU power usage in kilowatts
	CarbonIntensity float64 // Carbon intensity in gCO2/kWh
}

// EstimateCO2 estimates the CO₂ emissions based on CPU time and optional configuration.
// cpuSeconds: total CPU time in seconds.
// config: configuration with customizable parameters. If nil, defaults are used.
// Returns the estimated CO2 emissions in grams.
func EstimateCO2(cpuSeconds float64, config *Config) (float64, error) {
	if cpuSeconds <= 0 {
		return 0, errors.New("cpuSeconds must be greater than zero")
	}

	// Use defaults if no config is provided
	cpuPowerKW := defaultCPUPowerKW
	carbonIntensity := defaultCarbonIntensity
	if config != nil {
		if config.CPUPowerKW > 0 {
			cpuPowerKW = config.CPUPowerKW
		}
		if config.CarbonIntensity > 0 {
			carbonIntensity = config.CarbonIntensity
		}
	}

	// Step 1: Calculate energy consumption (kWh)
	timeHours := cpuSeconds / 3600.0
	energyConsumption := cpuPowerKW * timeHours

	// Step 2: Calculate CO₂ emissions (grams)
	co2 := energyConsumption * carbonIntensity

	return co2, nil
}

// Example usage function
func main() {
	cpuSeconds := 3600.0                     // 1 hour of CPU time
	co2, err := EstimateCO2(cpuSeconds, nil) // Use defaults
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Estimated CO2 emissions: %.2f grams\n", co2)

	// Custom configuration
	customConfig := &Config{
		CPUPowerKW:      0.08,  // 80W CPU
		CarbonIntensity: 300.0, // Lower carbon intensity
	}
	co2Custom, err := EstimateCO2(cpuSeconds, customConfig)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Estimated CO2 emissions (custom config): %.2f grams\n", co2Custom)
}
