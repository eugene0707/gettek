package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/eugene0707/gettek/models"
	"flag"
	"github.com/eugene0707/gettek/common"
)

func toJson(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func main() {
	common.CurrentDB().AutoMigrate(&models.Driver{}, &models.Metric{})

	fileType := flag.String("type", "drivers or metrics", "type of data in file to import")
	fileName := flag.String("file", "", "json file to import")
	flag.Parse()

	if *fileName == "" {
		fmt.Printf("Provide filename to import\n")
		os.Exit(1)
	}

	rawData := loadFile(*fileName)

	switch *fileType {
	case "drivers":
		importDrivers(rawData)
	case "metrics":
		importMetrics(rawData)
	default:
		fmt.Printf("Unknown type of data in file to import - %s.\n", *fileType)
		os.Exit(1)
	}
}

func loadFile(fn string) []byte {
	fmt.Printf("Loading data from %s.\n", fn)
	raw, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

func importDrivers(raw []byte) {
	fmt.Println("Importing drivers")

	var objArray []models.Driver
	err := json.Unmarshal(raw, &objArray)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, el := range objArray {
		common.CurrentDB().Save(&el)
		fmt.Println(toJson(el))
	}

}

func importMetrics(raw []byte) {
	fmt.Println("Importing metrics")

	var objArray []models.Metric
	err := json.Unmarshal(raw, &objArray)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, el := range objArray {
		common.CurrentDB().Save(&el)
		fmt.Println(toJson(el))
	}

}
