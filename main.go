package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RegionalConfig struct {
	url     string
	country string
}

// From: http://json2struct.mervine.net/?src=https://stat.ripe.net/data/country-resource-list/data.json?v4_format=prefix;resource=gb
type RegionalJSON struct {
	BuildVersion   string           `json:"build_version"`
	Cached         bool             `json:"cached"`
	Data           RegionalJSONData `json:"data"`
	DataCallStatus string           `json:"data_call_status"`
	Messages       [][]string       `json:"messages"`
	ProcessTime    int64            `json:"process_time"`
	QueryID        string           `json:"query_id"`
	SeeAlso        []interface{}    `json:"see_also"`
	ServerID       string           `json:"server_id"`
	Status         string           `json:"status"`
	StatusCode     int64            `json:"status_code"`
	Time           string           `json:"time"`
	Version        string           `json:"version"`
}

type RegionalJSONData struct {
	QueryTime string `json:"query_time"`
	Resources struct {
		Asn  []string `json:"asn"`
		Ipv4 []string `json:"ipv4"`
	} `json:"resources"`
}

func readFile(filename string) string {
	content, err := ioutil.ReadFile(filename) // the file is inside the local directory
	if err != nil {
		fmt.Printf("Err: Failed to read file %s", filename)
	}
	return string(content)
}

func readWeb(config RegionalConfig) string {
	resp, err := http.Get(fmt.Sprintf("%s=%s", config.url, config.country))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// TODO: if resp.StatusCode == http.StatusOK {
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body)

	return string(content)
}

func ToJSON(content string) RegionalJSON {
	var regionalJSON RegionalJSON
	// TODO: Test for invalid JSON
	json.Unmarshal([]byte(content), &regionalJSON)

	return regionalJSON
}

func main() {
	source := flag.String("source", "web", "Load data from the web or a local file")
	dataFile := flag.String("data", "example.json", "The filename to load data from")
	flag.Parse()
	config := RegionalConfig{"https://stat.ripe.net/data/country-resource-list/data.json?v4_format=prefix;resource", "gb"}

	var content string
	if *source == "file" {
		content = readFile(*dataFile)
	} else if *source == "web" {
		content = readWeb(config)
	} else {
		log.Fatal(fmt.Sprintf("Unrecognised source %s\n", *source))
	}

	json_content := ToJSON(content)
	//fmt.Printf("Build %s\n", json_content.BuildVersion)
	//fmt.Printf("QueryTime %s\n", json_content.Data.QueryTime)
	//fmt.Printf("IPv4 Addresses %s\n", json_content.Data.Resources.Ipv4[1])

	for _, ip_address := range json_content.Data.Resources.Ipv4 {
		fmt.Println(ip_address)
	}
}
