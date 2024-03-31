package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Configure Config

type Config struct {
	DbConfig struct {
		Database  string `yaml:"database"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Dbname    string `yaml:"dbname"`
		Charset   string `yaml:"charset"`
		ParseTime bool   `yaml:"parseTime"`
		Loc       string `yaml:"loc"`

		MaxIdleConns    int     `yaml:"maxIdleConns"`
		MaxOpenConns    int     `yaml:"maxOpenConns"`
		CoonMaxLifetime float64 `yaml:"coonMaxLifetime"`

		Path    string `yaml:"path"`
		Matched string `yaml:"matched"`
	} `yaml:"dbConfig"`

	Service struct {
		Port int `yaml:"port"`
	} `yaml:"service"`
}

func (c *Config) GetDbUrl() string {
	dbc := c.DbConfig
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s", dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.Dbname, dbc.Charset, dbc.ParseTime, dbc.Loc)
}

func init() {
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	} else if err = yaml.Unmarshal(file, &Configure); err != nil {
		log.Fatal(err)
	}
}
