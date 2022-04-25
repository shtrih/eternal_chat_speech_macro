package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PhraseOptions struct {
	Hotkeys []string
	Type    string
	Index   uint16
	Comment string
}

type Config struct {
	ChatHotkey string `mapstructure:"chat_hotkey"`
	Phrases    []PhraseOptions
}

var Types = []string{
	"attack",
	"flee",
	"follower",
	"hello",
	"hit",
	"idle",
	"intruder",
	"oppose",
	"service",
	"thief",
	"uniform",
	"ord_hello",
	"ord_intruder",
	"ord_idle",
	"ord_attack",
	"tb_idle",
	"tb_hello",
	"bm_idle",
	"bm_hello",
	"bm_flee",
	"bm_attack",
}

func (c *Config) FilterPhrases() {
	var filtered []PhraseOptions
	for _, option := range c.Phrases {
		if sliceContain(Types, option.Type) {
			filtered = append(filtered, option)
		} else {
			log.Warnf("Unknown type: '%s'.", option.Type)
		}
	}
	c.Phrases = filtered
}

func sliceContain(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func Init(configName string) *Config {
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./profiles")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: ", err)
	}

	C := new(Config)
	if err := viper.Unmarshal(C); err != nil {
		log.Fatal("unable to decode into struct: ", err)
	}
	C.FilterPhrases()
	return C
}
