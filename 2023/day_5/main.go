package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const (
	TYPE_SEED        = "seed"
	TYPE_SOIL        = "soil"
	TYPE_FERTILIZER  = "fertilizer"
	TYPE_WATER       = "water"
	TYPE_LIGHT       = "light"
	TYPE_TEMPERATURE = "temperature"
	TYPE_HUMIDITY    = "humidity"
	TYPE_LOCATION    = "location"
)

//go:embed input.txt
var input string

type Almanac struct {
	Seeds                 []Seed
	OriginalSeeds         []Seed
	SeedToSoil            []SourceToDestination
	SoilToFertilizer      []SourceToDestination
	FertilizerToWater     []SourceToDestination
	WaterToLight          []SourceToDestination
	LightToTemperature    []SourceToDestination
	TemperatureToHumidity []SourceToDestination
	HumidityToLocation    []SourceToDestination
}

type SourceToDestination struct {
	SourceStart      int
	DestinationStart int
	Range            int
}

type Seed struct {
	Key     int
	Touched bool
	Start   int
	End     int
}

func main() {
	fmt.Printf("Part 1: %d\n", solvePartOne(input))
	fmt.Printf("Part 2: %d\n", solvePartTwo(input))
}

func solvePartOne(input string) int {
	a := parseInput(input, false)
	minVal := 9999999999999999
	for _, ss := range a.OriginalSeeds {
		val := a.FindLocationValue(ss.Key, false)
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func solvePartTwo(input string) int {
	a := parseInput(input, true)
	minVal := 9999999999999999
	for _, ss := range a.OriginalSeeds {
		val := a.SolvePartTwo(ss, false)
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func (a *Almanac) SolvePartTwo(ss Seed, debug bool) int {
	minVal := 9999999999999999
	for v := ss.Start; v < ss.End; v++ {
		val := a.FindLocationValue(v, debug)
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func checkInRange(k int, start int, r int) bool {
	return k >= start && k < start+r
}

func findNextStep(rr []SourceToDestination, k int, debug bool, debugString string) int {
	for _, v := range rr {
		if checkInRange(k, v.SourceStart, v.Range) {
			diff := k - v.SourceStart
			k = v.DestinationStart + diff
			if debug {
				fmt.Printf("%s %d\n", debugString, k)
			}
			return k
		}
	}
	return k
}

func (a *Almanac) FindLocationValue(k int, debug bool) int {
	if debug {
		fmt.Printf("Seed %d\n", k)
	}
	k = findNextStep(a.SeedToSoil, k, debug, "Soil")
	k = findNextStep(a.SoilToFertilizer, k, debug, "Fertilizer")

	k = findNextStep(a.FertilizerToWater, k, debug, "Water")

	k = findNextStep(a.WaterToLight, k, debug, "Light")

	k = findNextStep(a.LightToTemperature, k, debug, "Temperature")

	k = findNextStep(a.TemperatureToHumidity, k, debug, "Humidity")

	k = findNextStep(a.HumidityToLocation, k, debug, "Location")

	return k
}

func (a *Almanac) ParseMap(idx int, splitted []string) int {
	ti := 0
	types := strings.Split(splitted[idx], "-")
	source := types[0]
	for i, row := range splitted[idx+1:] {
		if row == "" {
			return i
		}
		vv := strings.Split(row, " ")
		destinationStart, _ := strconv.Atoi(vv[0])
		sourceStart, _ := strconv.Atoi(vv[1])
		r, _ := strconv.Atoi(vv[2])

		a.SetSourceToDestination(source, sourceStart, destinationStart, r)

		ti++

	}
	return ti
}

func (a *Almanac) SetSourceToDestination(st string, ss int, ds int, r int) {
	switch st {
	case TYPE_SEED:
		a.SeedToSoil = append(a.SeedToSoil, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_SOIL:
		a.SoilToFertilizer = append(a.SoilToFertilizer, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_FERTILIZER:
		a.FertilizerToWater = append(a.FertilizerToWater, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_WATER:
		a.WaterToLight = append(a.WaterToLight, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_LIGHT:
		a.LightToTemperature = append(a.LightToTemperature, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_TEMPERATURE:
		a.TemperatureToHumidity = append(a.TemperatureToHumidity, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	case TYPE_HUMIDITY:
		a.HumidityToLocation = append(a.HumidityToLocation, SourceToDestination{SourceStart: ss, DestinationStart: ds, Range: r})
	}
}

func (a *Almanac) FindSeedInAlmanac(ss int) Seed {
	for _, v := range a.Seeds {
		if v.Key == ss {
			return v
		}
	}
	return Seed{Key: ss, Touched: true}
}

func parseInput(input string, useRange bool) Almanac {
	a := Almanac{}
	a.Seeds = []Seed{}
	a.OriginalSeeds = []Seed{}

	splitted := strings.Split(strings.TrimRight(input, "\n"), "\n")
	skip := 0
	for idx, row := range splitted {
		if skip > 0 {
			skip--
			continue
		}
		if row != "" {
			if strings.Contains(row, "seeds:") {
				seeds := parseSeeds(row, useRange)
				for _, s := range seeds {
					a.Seeds = append(a.Seeds, s)
					a.OriginalSeeds = append(a.OriginalSeeds, s)
				}
			} else if strings.Contains(row, "map:") {
				skip = a.ParseMap(idx, splitted)
			}
		}
	}
	return a
}

func parseSeeds(row string, useRange bool) []Seed {
	if useRange {
		return parseSeedsPartTwo(row)
	}
	ss := []Seed{}
	rowsplit := strings.Split(row, ":")[1]
	seeds := strings.Split(strings.Trim(rowsplit, " "), " ")
	for _, seed := range seeds {
		si, _ := strconv.Atoi(seed)
		s := Seed{Key: si}
		ss = append(ss, s)
	}
	return ss
}

func parseSeedsPartTwo(row string) []Seed {
	ss := []Seed{}
	rowsplit := strings.Split(row, ":")[1]
	seeds := strings.Split(strings.Trim(rowsplit, " "), " ")
	fs := 0
	// r := 0
	for idx, seed := range seeds {
		if idx%2 == 0 {
			fs, _ = strconv.Atoi(seed)
		} else {
			s, _ := strconv.Atoi(seed)
			ss = append(ss, Seed{Start: fs, End: fs + s})
		}
	}
	return ss
}
