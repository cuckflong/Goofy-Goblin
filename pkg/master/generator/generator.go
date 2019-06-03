// This package will provide functions to generate an agent binary with given modules, secret key, name and maybe other information

package generator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/johnathanclong/Goofy-Goblin/pkg/master/utils"
)

// GenerateAgent creates an agent with the given modules
func GenerateAgent(debug bool, verbose bool, silent bool, moduleList []string) error {
	tempDir, err := ioutil.TempDir(".", "compile")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	ldflags := "-X github.com/johnathanclong/Goofy-Goblin/pkg/agent/config.debugString=" + strconv.FormatBool(debug)
	ldflags += " -X github.com/johnathanclong/Goofy-Goblin/pkg/agent/config.verboseString=" + strconv.FormatBool(verbose)
	ldflags += " -X github.com/johnathanclong/Goofy-Goblin/pkg/agent/config.silentString=" + strconv.FormatBool(silent)

	var tags string

	for _, module := range moduleList {
		copyCompileFile("modules/"+module+"/exploit.go", tempDir)
		tag := getTag("modules/" + module + "/exploit.go")
		if tag != "" {
			if tags == "" {
				tags += tag
			} else {
				tags += " " + tag
			}
		}
	}

	copyCompileFile("cmd/goblin-agent/main.go", tempDir)

	cmd := exec.Command("go", "build", "-o", "../data/bin/test", "-ldflags", ldflags, "-tags", tags)
	cmd.Dir = tempDir
	cmd.Run()
	return err
}

func getTag(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	rp := regexp.MustCompile("// \\+build .+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if rp.MatchString(scanner.Text()) {
			tag := strings.Split(rp.FindString(scanner.Text()), " ")[2]
			return tag
		}
	}
	return ""
}

func copyCompileFile(src string, dir string) error {
	fileName, err := ioutil.TempFile(dir, "compile.*.go")
	if err != nil {
		utils.Status(utils.Error, fmt.Sprintf("Failed copying %s", src))
	}

	input, err := ioutil.ReadFile(src)

	if err != nil {
		utils.Status(utils.Error, fmt.Sprintf("Failed reading %s", src))
		return err
	}

	_, err = fileName.Write(input)
	return err
}
