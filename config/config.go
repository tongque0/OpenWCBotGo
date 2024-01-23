package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Env map[string]string

func InitEnv() {
	fmt.Print("Loading environment variables... ")
	Env = loadEnv()
	fmt.Println("Done.")
}

func loadEnv() map[string]string {
	// 读取 YAML 文件
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	// 解析 YAML
	var envMap map[string]string
	err = yaml.Unmarshal(data, &envMap)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
	}

	return envMap
}
