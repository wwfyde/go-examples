package main

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("docs")
	// Read config
	if err := viper.ReadInConfig(); err != nil {

	} else {

	}
	// Get value of Config items
	v1 := viper.Get("app.name")
	v2 := viper.GetString("app.name")
	fmt.Println(v1, reflect.TypeOf(v1))
	fmt.Println(v2, reflect.TypeOf(v2))
	fmt.Println()
	// Set config values
	viper.Set("app.name", "demo")

}
