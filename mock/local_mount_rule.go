package main

import (
	"amoeba"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

/*
  可用不同的方式来挂载规则，如数据库、配置中心、配置文件等，当前的读本地文件为其中一种
*/
func LocalMountRules() (map[string]*amoeba.Schema, error) {
	dir := "../rule/"
	files, _ := ioutil.ReadDir(dir)
	schemaMap := make(map[string]*amoeba.Schema, 0)
	for _, v := range files {
		fileName := v.Name()
		b, err := ioutil.ReadFile(dir + fileName)
		if err != nil {
			log.Panicf("failed to read the input file %+v", err)
			return schemaMap, err
		}
		var schema amoeba.Schema
		err = json.Unmarshal(b, &schema)
		if err != nil {
			log.Panicf("failed to unmarshal the input file %+v", err)
			return schemaMap, err
		}

		fileName = strings.TrimRight(fileName, ".json")
		schemaMap[fileName] = &schema
	}
	return schemaMap, nil
}
