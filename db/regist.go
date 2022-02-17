package db

import (
	"LibraryManagement/model"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var GlobalSettings *model.Settings

func InitRegist() {
	InitConfig()
	log.Println("========= init DB ===========")
	InitConn()
	log.Println("========= init DB success ===========")
}

func InitConfig() {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := fmt.Sprintf("%s%s", path, "/config/")
	datasource := model.DataSource{}
	cur := model.Settings{
		Datasource: datasource,
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("./config")
	readFile := "settings.yaml"
	viper.SetConfigName(readFile)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore errcode if desired
			log.Println("**** no such config file *****")
		} else {
			// Config file was found but another errcode was produced
			log.Println("****  read config errcode **** ")
			panic(err)
		}
		log.Fatal(err)
	}

	viper.Unmarshal(&cur)
	GlobalSettings = &cur
	log.Println("=========== settings init success ===========")
}
