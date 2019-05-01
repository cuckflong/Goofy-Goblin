// This package will provide information and functions for processing modules

package modules

import (
	"encoding/json"
	"fmt"
	"github.com/johnathanclong/Goofy-Goblin/pkg/master/utils"
	"io/ioutil"
	"os"
)

// Module Module struct containing the information of a module
type Module struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Platform    string   `json:"platform"`
	Options     []Option `json:"options"`
}

// Option Struct containing an option for a module
type Option struct {
	Index       int    `json:"index"`
	Name        string `json:"param"`
	Value       string `json:"value"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

// ParseInfo Parses the information of module
func ParseInfo(path string) {
	infoFile := "modules/" + path + "/info.json"
	jsonFile, err := os.Open(infoFile)
	if err != nil {
		utils.Status(utils.Error, fmt.Sprintf("Failed parsing %s", infoFile))
	}

	var moduleInfo Module
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &moduleInfo)
	utils.Status(utils.Success, moduleInfo.Name)
	utils.Status(utils.Success, moduleInfo.Description)
	utils.Status(utils.Success, moduleInfo.Path)
	utils.Status(utils.Success, moduleInfo.Platform)
	for _, option := range moduleInfo.Options {
		utils.Status(utils.Success, option.Name)
		utils.Status(utils.Success, option.Description)
	}
}
