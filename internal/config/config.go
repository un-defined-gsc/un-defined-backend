package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	Http     Http     `yaml:"http"`
	Log      Log      `yaml:"log"`
	App      App      `yaml:"app"`
	Email    Email    `yaml:"email"`
	Data     Data     `yaml:"data"`
}

type Data struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Email struct {
	Address  string `yaml:"address"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Database struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Database       string `yaml:"database"`
	MigrationsPath string `yaml:"migrations_path"`
}

type Http struct {
	Port             string   `yaml:"port"`
	ReadTimeout      string   `yaml:"readTimeout"`
	WriteTimeout     string   `yaml:"writeTimeout"`
	AllowedOrigins   []string `yaml:"allowedOrigins"`
	AllowedMethods   []string `yaml:"allowedMethods"`
	AllowedHeaders   []string `yaml:"allowedHeaders"`
	ExposedHeaders   []string `yaml:"exposedHeaders"`
	MaxAge           int      `yaml:"maxAge"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	ProxyHeader      string   `yaml:"proxyHeader"`
}

type Log struct {
	Level string `yaml:"log_level"`
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Site    string `yaml:"site"`
	Prod    *bool  `yaml:"prod"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

var (
	CONFIG_PATH = "./config.yaml"
	config      Config
)

func InitializeConfig(configPath ...string) (err error) {
	if len(configPath) > 0 {
		CONFIG_PATH = configPath[0]
	}
	config = Config{}
	return config.load()
}

func (c *Config) load() (err error) {
	file, err := validateConfigPath()
	if err != nil {
		return err
	}
	defer file.Close()
	yaml.NewDecoder(file).Decode(c)
	return nil
}

func (c *Config) DatabaseConfig() Database {
	return config.Database
}

func validateConfigPath() (file *os.File, err error) {
	s, err := os.Stat(CONFIG_PATH)
	if err != nil {
		return
	}
	if s.IsDir() {
		err = fmt.Errorf("'%s' is a directory, not config file", CONFIG_PATH)
		return
	}
	file, err = os.Open(CONFIG_PATH)
	return
}

func GetConfig() Config {
	return config
}
