package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type App struct {
	Port        int    `yaml:"port"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

type Database struct {
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	DbName       string `yaml:"db_name"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Charset      string `yaml:"charset"`
	ParseTime    bool   `yaml:"parse_time"`
	Loc          string `yaml:"loc"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Jwt struct {
	Secret    string `yaml:"secret"`
	ExpiresIn string `yaml:"expires_in"`
}

type RefreshToken struct {
	Secret    string `yaml:"secret"`
	ExpiresIn string `yaml:"expires_in"`
}

type Config struct {
	Env          string       `yaml:"env"`
	App          App          `yaml:"app"`
	Database     Database     `yaml:"database"`
	Redis        Redis        `yaml:"redis"`
	Jwt          Jwt          `yaml:"jwt"`
	RefreshToken RefreshToken `yaml:"refresh_token"`
}

func AppConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flagPath := flag.String("config", "", "Path to config file (YAML)")
		flag.Parse()
		configPath = *flagPath
	}

	if configPath == "" {
		configPath = "config/local.yml"
		fmt.Println("[config] No config path provided, using default:", configPath)
	}

	//-------------------------------------------
	//----------------- Note --------------------
	// %q : when you want to see exactly how the string
	// is encoded (including quotes and escapes).
	//
	// %w : Only valid inside the format string passed to fmt.Errorf
	//-------------------------------------------

	// Read YAML file
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("unable to read config file %q: %w", configPath, err))
	}

	// Parse YAML into Config struct
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("failed to parse YAML in %q: %w", configPath, err))
	}

	return &cfg, nil
}
