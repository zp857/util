package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func New(filename string, val interface{}) (err error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&val)
	return
}

func NewWithWatch(filename string, val interface{}) (err error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&val)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed", e.Name)
	})
	return
}
