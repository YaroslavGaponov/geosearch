package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/YaroslavGaponov/geosearch"
)

const (
	Width = 50

	ZIP_FILE = "./worldcities.zip"
	CSV_FILE = "worldcities.csv"

	attemps                   = 5
	distanceBetweenNeighbours = 400
)

func main() {

	zipFile, err := zip.OpenReader(ZIP_FILE)
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}
	defer zipFile.Close()

	var file *zip.File
	for _, f := range zipFile.File {
		if f.Name == CSV_FILE {
			file = f
			break
		}
	}
	if file == nil {
		log.Fatalf("file %s is not found", CSV_FILE)
	}

	reader, err := file.Open()
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}
	defer reader.Close()

	gs := geosearch.New(attemps, distanceBetweenNeighbours)

	fmt.Println("Creating geo structure...")
	counter := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		city := trim(parts[1])
		lat, err := strconv.ParseFloat(trim(parts[2]), 64)
		if err != nil {
			continue
		}
		lng, err := strconv.ParseFloat(trim(parts[3]), 64)
		if err != nil {
			continue
		}
		print(city)
		gs.AddObject(&geosearch.Object{Id: city, Point: geosearch.Point{Latitude: lat, Longitude: lng}})
		counter++
	}
	printEnd("Done")
	fmt.Println(counter, "points are loaded")

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}

	result := gs.Search(me)
	fmt.Println("Result", result.Object.Id, result.Distance, "km")
}

func trim(s string) string {
	return s[1 : len(s)-1]
}

func print(s string) {
	fmt.Print(s, strings.Repeat(" ", Width-len(s)), "\r")
}
func printEnd(s string) {
	fmt.Println(s, strings.Repeat(" ", Width-len(s)))
}
