package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Node struct {
	Test     string `json:"test"`
	Children []Node `json:"children"`
}

var stRootDir02 string
var stSeparator02 string
var stJsonFileName02 = "dir.json"
var iRootNode Node

func loadJson02() {
	stSeparator02 = string(filepath.Separator)
	stWorkDir, _ := os.Getwd()

	stRootDir02 = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator02)]
	gnJsonFileBytes, _ := os.ReadFile(stWorkDir + stSeparator02 + stJsonFileName02)
	if json.Unmarshal(gnJsonFileBytes, &iRootNode) != nil {
		panic("Load Json Data Error" + json.Unmarshal(gnJsonFileBytes, &iRootNode).Error())
	}

}
func TestGenerateDir02(t *testing.T) {
	loadJson02()
}
