// This package will provide functions for the master server

package master

import (
	"encoding/json"
	"fmt"
	"github.com/johnathanclong/Goofy-Goblin/pkg/utils"
	"io/ioutil"
	"os"
)

// ModuleInfo Module struct containing the information of a module
type ModuleInfo struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Path        string         `json:"path"`
	Platform    string         `json:"platform"`
	Options     []ModuleOption `json:"options"`
}

// ModuleOption Struct containing an option for a module
type ModuleOption struct {
	Index  int    `json:"index"`
	Param  string `json:"param"`
	Detail string `json:"detail"`
}

// ParseInfo Parses the information of module
func ParseInfo(path string) {
	infoFile := "modules/" + path + "/info.json"
	jsonFile, err := os.Open(infoFile)
	if err != nil {
		utils.Status(utils.Error, fmt.Sprintf("Failed parsing %s", infoFile))
	}

	var moduleInfo ModuleInfo
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &moduleInfo)
	utils.Status(utils.Success, moduleInfo.Name)
	utils.Status(utils.Success, moduleInfo.Description)
	utils.Status(utils.Success, moduleInfo.Path)
	utils.Status(utils.Success, moduleInfo.Platform)
	for _, option := range moduleInfo.Options {
		utils.Status(utils.Success, option.Param)
		utils.Status(utils.Success, option.Detail)
	}
}
