package variablecollector

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/* variable collector
collecting variables need to follow this order
1- env variable is the first thing and overrides all bellow methodes
2- collecting the values from a specific file when providing a file path 
3- if all of the folowing is not present an error should be displayed  also we can provide some default values for port and protocol
*/

type variableInfo struct {
	key, value string
	isDefined  bool
}

var variableInfoList []variableInfo

// variable info constructor
func newVariableInfoConstructor(variableKey string) *variableInfo {
	return &variableInfo{
		key: variableKey,
	}
}

// variable info list constructor
func NewVariableListConstructor(variableKeys ...string) {
	for _, variableKey := range variableKeys {
		variableInfo := newVariableInfoConstructor(variableKey)
		variableInfoList = append(variableInfoList, *variableInfo)
	}
}

func populateEnvValues() {
	fmt.Println("searching values from operating system if exists")
	for index := range variableInfoList {
		variableInfoPtr := &variableInfoList[index]

		variableInfoPtr.value = os.Getenv(variableInfoPtr.key)
		variableInfoPtr.isDefined = variableInfoPtr.value != ""

	}
}

func populateFileValues(envFilePath string) {
	_, err := os.Stat(envFilePath)
	if os.IsNotExist(err) {
		fmt.Println("file with path: ", envFilePath, " does not exists")
	}
	fmt.Println("searching values from env file")
	for index := range variableInfoList {
		variableInfoPtr := &variableInfoList[index]
		if !variableInfoPtr.isDefined {
			err := fetchFromFile(envFilePath, variableInfoPtr)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func fetchFromFile(envFilePath string, varinfo *variableInfo) error {
	file, err := os.Open(envFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) == 2 && strings.ToUpper(parts[0]) == strings.ToUpper(varinfo.key) {
			varinfo.value = parts[1]
			varinfo.isDefined = true
		}
	}

	return err
}

func ExtractVariableValues(envFilePath string )  map[string]string {
	//testing time execution 
	now := time.Now()
	VariableValueMap := make(map[string]string)
	populateEnvValues()
	populateFileValues(envFilePath)
	for _, variableInfoObject := range variableInfoList {
		if variableInfoObject.isDefined {
			VariableValueMap[variableInfoObject.key] = variableInfoObject.value
		}
	}
	fmt.Println("took time: ", time.Since(now))
	return VariableValueMap
}
