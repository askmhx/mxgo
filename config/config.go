package config

import (
	"os"
	"strconv"
	"bufio"
	"strings"
)

type Config struct {
	data map[string]map[string]string
}

func NewConfig(configPath string) *Config {
	config := &Config{}
	config.Load(configPath)
	return config
}

func (config *Config) String(key string) string {
	for _,kv := range config.data {
		for k, v := range kv {
			if k == key {
				return v
			}
		}
	}
	return ""
}

func (config *Config) Int(key string) int {
	value , _ := strconv.Atoi(config.String(key))
	return value
}

func (config *Config) Bool(key string) bool {
	value , _ := strconv.Atoi(config.String(key))
	return value!=0
}


func (config *Config) Load(configPath string) {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	config.data = make(map[string]map[string]string)
	sectionTitle := ""
	section := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				if len(sectionTitle)>0 {
					config.data[sectionTitle] = section
				}
				sectionTitle = line[1:len(line)-1]
				section = make(map[string]string)
			}else if strings.Contains(line, "=") {
				lineArray := strings.Split(line, "=")
				if len(lineArray) != 2 {
					panic("config format error:" + configPath + "\n at:" + line)
				}
				key := lineArray[0]
				value := lineArray[1]
				if len(sectionTitle)==0 {
					sectionTitle = "default"
				}
				section[key] = value
			}
		}
	}
	if  len(sectionTitle)>0 {
		config.data[sectionTitle] = section
	}
}

