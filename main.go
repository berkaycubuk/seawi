package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type Data struct {
	FirstName string `yaml:"FirstName"`
	LastName string `yaml:"LastName"`
}

func getConfiguration(path string) (data *Data, error error) {
	// read configuration file
	rawData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// convert byte array to meaningful Data struct
	var convertedData Data
	if err = yaml.Unmarshal(rawData, &convertedData); err != nil {
		return nil, err
	}

	return &convertedData, nil
}

func main() {
	fmt.Println("Seawi v1.0.0")

	configuration, err := getConfiguration("./data.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	// template
	t, err := template.ParseFiles("template/base.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create output file and write parsed template into it
	f, err := os.Create("./output.html")
	w := bufio.NewWriter(f)

	t.Execute(w, configuration)

	w.Flush()

	fmt.Println("Done")
}
