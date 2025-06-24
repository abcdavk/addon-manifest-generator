package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Modules struct {
	Type    string `json:"type"`
	UUID    string `json:"uuid"`
	Version [3]int `json:"version"`
}

type Dependencies struct {
	UUID    string `json:"uuid"`
	Version [3]int `json:"version"`
}

type ManifestBehavior struct {
	FormatVersion int `json:"format_version"`
	Header        struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		UUID         string `json:"uuid"`
		Version      [3]int `json:"version"`
		MinEngineVer [3]int `json:"min_engine_version"`
	} `json:"header"`
	Modules      []Modules      `json:"modules"`
	Dependencies []Dependencies `json:"dependencies"`
}

func main() {
	mnbp := ManifestBehavior{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Addon Name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if len(input) != 0 {
			mnbp.Header.Name = input
			break
		} else {
			mnbp.Header.Name = "Addon Template"
			break
		}
	}

	for {
		fmt.Print("Addon Description: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if len(input) != 0 {
			mnbp.Header.Description = input
			break
		} else {
			mnbp.Header.Description = "Cool addon temolate!"
			break
		}
	}

	mnbp.FormatVersion = 2

	mnbp.Header.UUID = uuid.NewString()
	mnbp.Header.Version = [3]int{1, 0, 0}
	mnbp.Header.MinEngineVer = [3]int{1, 21, 0}
	mnbp.Modules = append(mnbp.Modules, Modules{
		Type:    "data",
		UUID:    uuid.NewString(),
		Version: [3]int{1, 0, 0},
	})
	mnbp.Dependencies = append(mnbp.Dependencies, Dependencies{
		UUID:    uuid.NewString(),
		Version: [3]int{1, 0, 0},
	})

	output, err := json.MarshalIndent(mnbp, "", "	")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("BP/manifest.json", []byte(output), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
