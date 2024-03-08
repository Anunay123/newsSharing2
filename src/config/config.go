package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

func DoInit() {

	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath("settings/")
	if err := Config.ReadInConfig(); err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}
	//TO Watch config; If not this statement, changes will not reflect.
	Config.WatchConfig()
	Config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

}
