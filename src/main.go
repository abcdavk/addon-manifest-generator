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

type ModuleTypeData struct {
	Type    string `json:"type"`
	UUID    string `json:"uuid"`
	Version [3]int `json:"version"`
}

type ModuleTypeScript struct {
	Type        string `json:"type"`
	UUID        string `json:"uuid"`
	Version     [3]int `json:"version"`
	Description string `json:"description"`
	Lang        string `json:"language"`
	Entry       string `json:"entry"`
}

type DependencyAddon struct {
	UUID    string `json:"uuid"`
	Version [3]int `json:"version"`
}

type DependencyScript struct {
	ModuleName string `json:"module_name"`
	Version    string `json:"version"`
}

type Header struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	UUID         string `json:"uuid"`
	Version      [3]int `json:"version"`
	MinEngineVer [3]int `json:"min_engine_version"`
}

type ManifestBehavior struct {
	FormatVersion int    `json:"format_version"`
	Header        Header `json:"header"`
	Modules       []any  `json:"modules"`
	Dependencies  []any  `json:"dependencies"`
}

type ManifestResources struct {
	FormatVersion int    `json:"format_version"`
	Header        Header `json:"header"`
	Modules       []any  `json:"modules"`
	Dependencies  []any  `json:"dependencies"`
}

func main() {
	mnbp := ManifestBehavior{}
	mnrp := ManifestResources{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Addon Name: ")
	addonNameInput, _ := reader.ReadString('\n')
	addonNameInput = strings.TrimSuffix(addonNameInput, "\n")
	if len(addonNameInput) != 0 {
		mnbp.Header.Name = addonNameInput
		mnrp.Header.Name = addonNameInput
	} else {
		mnbp.Header.Name = "This is addon name"
		mnrp.Header.Name = "This is addon name"
	}

	fmt.Print("Addon Description: ")
	addonDescInput, _ := reader.ReadString('\n')
	addonDescInput = strings.TrimSuffix(addonDescInput, "\n")
	if len(addonDescInput) != 0 {
		mnbp.Header.Description = addonDescInput
		mnrp.Header.Description = addonDescInput
	} else {
		mnbp.Header.Description = "This is description!"
		mnrp.Header.Description = "This is description!"
	}

	fmt.Print("Use script API? (yes/no): ")
	useScriptInput, _ := reader.ReadString('\n')
	useScriptInput = strings.TrimSuffix(useScriptInput, "\n")
	if useScriptInput == "no" {

	} else {
		moduleName := "@minecraft/server"
		moduleVer := conn(moduleName)

		fmt.Printf("  @minecraft/server version (%s): ", moduleVer)
		scriptVersionInput, _ := reader.ReadString('\n')
		scriptVersionInput = strings.TrimSuffix(scriptVersionInput, "\n")

		scriptDepend := DependencyScript{
			ModuleName: moduleName,
			Version:    scriptVersionInput,
		}

		if scriptVersionInput == "" {
			scriptDepend.Version = moduleVer
		}

		mnbp.Dependencies = append(mnbp.Dependencies, scriptDepend)

		moduleName = "@minecraft/server-ui"
		moduleVer = conn(moduleName)

		fmt.Printf("  @minecraft/server-ui version (default: %s): ", moduleVer)
		scriptUIVersionInput, _ := reader.ReadString('\n')
		scriptUIVersionInput = strings.TrimSuffix(scriptUIVersionInput, "\n")

		scriptUIDepend := DependencyScript{
			ModuleName: "@minecraft/server-ui",
			Version:    scriptUIVersionInput,
		}

		if scriptUIVersionInput == "" {
			scriptUIDepend.Version = moduleVer
		}

		mnbp.Dependencies = append(mnbp.Dependencies, scriptUIDepend)

		mnbp.Modules = append(mnbp.Modules, ModuleTypeScript{
			Type:        "script",
			UUID:        uuid.NewString(),
			Version:     [3]int{1, 0, 0},
			Description: "Script Resource",
			Lang:        "javascript",
			Entry:       "scripts/main.js",
		})
	}

	mnbp.FormatVersion = 2
	mnrp.FormatVersion = 2

	mnbp.Header.UUID = uuid.NewString()
	mnrp.Header.UUID = uuid.NewString()

	mnbp.Header.Version = [3]int{1, 0, 0}
	mnbp.Header.MinEngineVer = [3]int{1, 21, 0}

	mnrp.Header.Version = [3]int{1, 0, 0}
	mnrp.Header.MinEngineVer = [3]int{1, 21, 0}

	mnbp.Modules = append(mnbp.Modules, ModuleTypeData{
		Type:    "data",
		UUID:    uuid.NewString(),
		Version: [3]int{1, 0, 0},
	})

	mnrp.Modules = append(mnrp.Modules, ModuleTypeData{
		Type:    "resources",
		UUID:    uuid.NewString(),
		Version: [3]int{1, 0, 0},
	})

	mnbp.Dependencies = append(mnbp.Dependencies, DependencyAddon{
		UUID:    mnrp.Header.UUID,
		Version: [3]int{1, 0, 0},
	})

	mnrp.Dependencies = append(mnrp.Dependencies, DependencyAddon{
		UUID:    mnbp.Header.UUID,
		Version: [3]int{1, 0, 0},
	})

	output, err := json.MarshalIndent(mnbp, "", "	")

	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("packs", os.ModePerm)
	os.Mkdir("packs/BP", os.ModePerm)
	os.Mkdir("packs/RP", os.ModePerm)

	err = os.WriteFile("packs/BP/manifest.json", []byte(output), 0666)
	if err != nil {
		log.Fatal(err)
	}

	output, err = json.MarshalIndent(mnrp, "", "	")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("packs/RP/manifest.json", []byte(output), 0666)
	if err != nil {
		log.Fatal(err)
	}

	println("Successfully created manifests!")
	println("packs/BP/manifest.json")
	println("packs/RP/manifest.json")
}
