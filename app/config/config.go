package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Database struct {
	Username        string `yaml:"username"`
	Pwd             string `yaml:"pwd"`
	DriverName      string `yaml:"driver_name"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	DatabaseName    string `yaml:"database_name"`
	TableSchemaName string `yaml:"table_schema_name"`
	ShowSQL         bool   `yaml:"show_sql"`
	Sync            bool   `yaml:"sync"`
}

type Service struct {
	Port      string `yaml:"port"`
	UploadDir string `yaml:"upload_dir"`
}

type Log struct {
	LogFilePath string `yaml:"log_file_path"`
	LogFileName string `yaml:"log_file_name"`
	Day         int    `yaml:"day"`
}

type Config struct {
	Database     Database `yaml:"database"`
	Service      Service  `yaml:"service"`
	Log          Log      `yaml:"logger"`
	FrontendPath string   `yaml:"frontend_path"`
}

func LoadConfig(filepath string) (*Config, error) {
	configFilePath := filepath
	if configFilePath == "" {
		configFilePath = "./config.yaml"
	}
	if len(os.Args) > 1 {
		configFilePath = os.Args[1]
	}
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
