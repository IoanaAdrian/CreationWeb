package utilities

import (
	"io/ioutil"
	"strings"
)

var MAIN_DIR_PATH = ""
var URL = ""
var PASSOWRD = ""

func SetPassword(s string){
	PASSOWRD = s
}
func GetPassword() string{
	return PASSOWRD
}

func SetMainDirPath(s string) {
	MAIN_DIR_PATH = s
}

func GetMainDirPath() string {
	return MAIN_DIR_PATH
}
func GetURL() string {
	return URL
}
func SetURL(s string) {
	URL = s
}

func SetConfig() {
	b, err := ioutil.ReadFile("./config.txt")
	HandleErr(err)
	splitConfig := strings.Split(string(b), "\n")
	for _, s := range splitConfig {
		t := strings.Split(s, " ")
		switch t[0] {
		case "PASS":
			SetPassword(t[1])
			break
		case "PATH":
			SetMainDirPath(t[1])
			break
		case "URL":
			SetURL(t[1])
			break
		}
	}
}
