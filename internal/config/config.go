package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
	conf "github.com/Jofich/Blog-website/config"
)

type Config struct {
	ServerCfg `yaml:"server"`
	DBCfg     `yaml:"database"`
}

type DBCfg struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	DB_name  string `yaml:"db_name"`
}

type ServerCfg struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Load() *Config {

	cfgPath := "config/config.yaml"

	if(len(conf.JwtKey) == 0){
		log.Fatalln("JwtKey not found in config/JwtKey.go")
	}

	cfgFile, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatalf("Couldn open file %s: %v", cfgPath, err)
	}
	var cfg Config
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		log.Fatalf("Coulnd unmarshall file : %v", err)
	}

	return &cfg
}
