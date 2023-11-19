package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func main() {
	if err := decode(); err != nil {
		fmt.Println("Error:", err)
	}

}

type coordResult struct {
	point orb.Point
	err   error
}

// func convertPostList(postList string) (orb.LineString, error) {
// 	coords := strings.Split(postList, " ")
// 	results := make([]orb.Point, len(coords)/2)
// 	var wg sync.WaitGroup

// 	for i := 0; i < len(coords)/2; i++ {
// 		first := coords[i*2]
// 		second := coords[i*2+1]
// 		wg.Add(1)

// 		go func(i int, first, second string) {
// 			defer wg.Done()
// 			point, err := processCoord(first, second)
// 			if err != nil {
// 				// Handle error as needed
// 				return
// 			}
// 			results[i] = point
// 		}(i, first, second)
// 	}

// 	wg.Wait()

// 	return orb.LineString(results), nil
// }

func convertPostList(postList string) (orb.LineString, error) {
	coords := strings.Split(postList, " ")
	results := make([]orb.Point, len(coords)/2)
	var wg sync.WaitGroup

	for i := 0; i < len(coords)/2; i++ {
		first := coords[i*2]
		second := coords[i*2+1]
		wg.Add(1)

		go func(i int, first, second string) {
			defer wg.Done()
			point, err := processCoord(first, second)
			if err != nil {
				// Handle error as needed
				return
			}
			results[i] = point
		}(i, first, second)
	}

	wg.Wait()

	return orb.LineString(results), nil
}

func processCoord(first, second string) (orb.Point, error) {
	firstCoord, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return orb.Point{}, fmt.Errorf("error parsing first coordinate: %v", err)
	}

	secondCoord, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return orb.Point{}, fmt.Errorf("error parsing second coordinate: %v", err)
	}

	return orb.Point{secondCoord, firstCoord}, nil
}

func decode() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}

	// Construct the path to the XML file
	xmlFilePath := filepath.Join(cwd, "RailTransportNetwork.gml")

	// Open the XML file
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		fmt.Println("Error opening XML file:", err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	fc := geojson.NewFeatureCollection()

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error decoding XML token: %v", err)
		}
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "RailwayLink" {
				var rl RailwayLink
				err := decoder.DecodeElement(&rl, &se)
				if err != nil {
					return fmt.Errorf("error decoding RailwayLink element: %v", err)
				}

				// geometry, err := convertPostList(rl.CentrelineGeometry.LineString.PosList)
				geometry, err := convertPostList(rl.CentrelineGeometry.LineString.PosList)
				if err != nil {
					return fmt.Errorf("error converting coordinate list: %v", err)
				}

				props := geojson.Properties{
					"ID":                   rl.ID,
					"BeginLifespanVersion": rl.BeginLifespanVersion,
					"EndLifespanVersion":   rl.EndLifespanVersion.Text,
					"Fictitious":           rl.Fictitious,
				}

				ft := geojson.Feature{
					Type:       "Feature",
					Geometry:   geometry,
					Properties: props,
				}

				fc.Append(&ft)
			}
		}
	}

	// Marshal the feature collection to GeoJSON
	geoJSON, err := fc.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error marshaling to GeoJSON: %v", err)
	}

	// Save the GeoJSON data to a file using ioutil.WriteFile
	err = ioutil.WriteFile("output.geojson", geoJSON, 0644)
	if err != nil {
		return fmt.Errorf("error saving GeoJSON to file: %v", err)
	}

	return nil
}
