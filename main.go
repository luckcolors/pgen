package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"encoding/json"
)

type Config struct {
	NumberOfPositions uint
	Dictionaries      map[uint]map[uint]string
}

var CFG = &Config{}

func MakeTemplateMap() map[uint]map[uint]string {
	var TemplateMap = make(map[uint]map[uint]string)
	TemplateMap[0] = make(map[uint]string)
	TemplateMap[1] = make(map[uint]string)
	TemplateMap[2] = make(map[uint]string)
	TemplateMap[3] = make(map[uint]string)
	TemplateMap[0][0] = "a"
	TemplateMap[0][1] = "b"
	TemplateMap[0][2] = "c"
	TemplateMap[1][0] = "a"
	TemplateMap[1][1] = "f"
	TemplateMap[2][0] = "g"
	TemplateMap[2][1] = "h"
	TemplateMap[3][0] = "ffffff"
	return TemplateMap
}

func HandleCfg() {
	if len(os.Args) <= 1 {
		fmt.Println("pgen <command> <filename> \n", "   makecfg: creates a default cfg file1\n", "   run: runs the program with the specified\n", "Author: luckcolors")
		os.Exit(0)
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "makecfg" {
			_, err := os.Stat(os.Args[2])
			if err != nil {
				tmp, err := json.Marshal(Config{
					NumberOfPositions: 4,
					Dictionaries:      MakeTemplateMap(),
				})
				err = ioutil.WriteFile(os.Args[2], tmp, 777)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Created template config file: " + os.Args[2])
				os.Exit(2)
			}
			fmt.Println("Template config file already exists: " + os.Args[2])
			os.Exit(1)
		}
		if os.Args[1] == "run" {
			cfgData, err := ioutil.ReadFile(os.Args[2])
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(cfgData, CFG)
			if err != nil {
				panic(err)
			}
		}
	} else {
		fmt.Println("pgen: invalid command/arguments")
		os.Exit(1)
	}
	return
}

func CGen(position uint, buffer []string) {
	if position == CFG.NumberOfPositions {
		position--
		for _, word := range CFG.Dictionaries[position] {
			buffer[position] = word
			fmt.Println(ArrayToString(buffer))
		}
	} else {
		for _, word := range CFG.Dictionaries[position] {
			buffer[position] = word
			CGen((position + 1), buffer)
		}
	}
}

func ArrayToString(array []string) (out string) {
	for _, data := range array {
		out = out + data
	}
	return out
}

func main() {
	HandleCfg()
	var buffer = make([]string, CFG.NumberOfPositions) // 0 1 2
	CGen(0, buffer)
}
