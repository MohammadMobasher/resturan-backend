package config

import (
	"fmt"

	"github.com/MohammadMobasher/resturan-backend/models"
	"github.com/spf13/viper"
)

func GetConfig() models.Configuration {
	config := models.Configuration{}

	viper.SetConfigName("")    // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")  // path to look for the config file in
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(err)
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil { // Handle errors reading the config file
		panic(err)

	}

	return config
}
