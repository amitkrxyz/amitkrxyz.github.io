package main

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Link struct {
	Name string
	Url  string
}

func LinksToString(links []Link) (names []string, urls []string) {
	names = make([]string, len(links))
	urls = make([]string, len(links))
	for i, l := range links {
		names[i] = l.Name
		urls[i] = l.Url
	}
	return names, urls
}

type Config struct {
	Username string `toml:"username"`
	Url      string `toml:"url"`
	Source   string `toml:"source"`
	AsciiArt string `toml:"ascii_art"`
	About    string `toml:"about"`
	Links    []Link `toml:"links"`
}

func NewConfig(configFileName string) (Config, error) {
	file, err := os.ReadFile(configFileName)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = toml.Unmarshal(file, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
