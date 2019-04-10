// This package will provide functions to generate an agent binary with given modules, secret key, name and maybe other information

package generator

import (
	"io/ioutil"
	"os"
	"strconv"
)

// GenerateAgent creates an agent with the given modules
func GenerateAgent(debug bool, verbose bool, silent bool, moduleList []string) error {
	tempDir, err := ioutil.TempDir("", "compile")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	ldflags := "-X github.com/johnathanclong/Goofy-Goblin/pkg/config.debugString=" + strconv.FormatBool(debug)
	ldflags += "-X github.com/johnathanclong/Goofy-Goblin/pkg/config.verboseString=" + strconv.FormatBool(verbose)
	ldflags += "-X github.com/johnathanclong/Goofy-Goblin/pkg/config.silentString=" + strconv.FormatBool(silent)

	var modules []string

	for i, module := range moduleList {
		if i == 0 {
			modules = append(modules, module)
		} else {
			modules = append(modules, " "+module)
		}
	}

	return err
}
