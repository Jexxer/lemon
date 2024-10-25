package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var volumeConversions = map[string]float64{
	"cups_to_milliliters":        236.588,
	"milliliters_to_cups":        1 / 236.588,
	"cups_to_liters":             0.236588,
	"liters_to_cups":             1 / 0.236588,
	"tablespoons_to_milliliters": 14.7868,
	"milliliters_to_tablespoons": 1 / 14.7868,
	"teaspoons_to_milliliters":   4.92892,
	"milliliters_to_teaspoons":   1 / 4.92892,
	"tablespoons_to_teaspoons":   3.0,
	"teaspoons_to_tablespoons":   1 / 3.0,
	"cups_to_tablespoons":        16.0,
	"tablespoons_to_cups":        1 / 16.0,
	"pints_to_cups":              2.0,
	"cups_to_pints":              1 / 2.0,
	"quarts_to_cups":             4.0,
	"cups_to_quarts":             1 / 4.0,
	"gallons_to_cups":            16.0,
	"cups_to_gallons":            1 / 16.0,
	"liters_to_milliliters":      1000.0,
	"milliliters_to_liters":      1 / 1000.0,
}

var weightConversions = map[string]float64{
	"grams_to_ounces":     0.03527396,
	"ounces_to_grams":     1 / 0.03527396,
	"grams_to_pounds":     0.00220462,
	"pounds_to_grams":     1 / 0.00220462,
	"ounces_to_pounds":    0.0625,
	"pounds_to_ounces":    16.0,
	"kilograms_to_pounds": 2.20462,
	"pounds_to_kilograms": 1 / 2.20462,
	"grams_to_kilograms":  1 / 1000.0,
	"kilograms_to_grams":  1000.0,
}

var volumeUnits = map[string]bool{
	"cups": true, "milliliters": true, "liters": true, "tablespoons": true, "teaspoons": true, "pints": true, "quarts": true, "gallons": true,
}

var weightUnits = map[string]bool{
	"grams": true, "ounces": true, "pounds": true, "kilograms": true,
}

func convertMeasurement(value float64, unitFrom string, unitTo string, conversionMap map[string]float64) (float64, error) {
	key := fmt.Sprintf("%s_to_%s", strings.ToLower(unitFrom), strings.ToLower(unitTo))
	conversionFactor, exists := conversionMap[key]
	if !exists {
		return 0, fmt.Errorf("conversion from %s to %s not supported", unitFrom, unitTo)
	}
	return value * conversionFactor, nil
}

func main() {
	var value float64
	var unitFrom, unitTo string

	flag.Float64Var(&value, "value", 0.0, "Measurement value to convert")
	flag.StringVar(&unitFrom, "from", "", "Unit to convert from")
	flag.StringVar(&unitTo, "to", "", "Unit to convert to")
	flag.Parse()

	unitFromLower := strings.ToLower(unitFrom)
	unitToLower := strings.ToLower(unitTo)

	if volumeUnits[unitFromLower] && volumeUnits[unitToLower] {
		result, err := convertMeasurement(value, unitFromLower, unitToLower, volumeConversions)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f %s is %.2f %s\n", value, unitFrom, result, unitTo)
	} else if weightUnits[unitFromLower] && weightUnits[unitToLower] {
		result, err := convertMeasurement(value, unitFromLower, unitToLower, weightConversions)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f %s is %.2f %s\n", value, unitFrom, result, unitTo)
	} else {
		log.Fatal("Invalid unit conversion. Make sure you're using valid units for volume or weight.")
	}
}
