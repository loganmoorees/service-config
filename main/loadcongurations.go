package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LoadConfigurations interface {
	LoadConfig() *Server
}

type Configuration struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Server struct {
	Server Configuration `yaml:"server"`
}

func LoadConfig() *Server {
	file, _ := ioutil.ReadFile("config/config.yml")
	fmt.Println(file)
	server := new(Server)
	err := yaml.Unmarshal(file, server)
	if err != nil {
		return nil
	}
	return server
}
