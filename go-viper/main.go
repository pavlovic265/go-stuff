package main

import (
	"fmt"

	"example.com/goviper/util"
)

func main() {

	// vp := viper.New()

	// // test.json
	// vp.SetConfigName("test")
	// vp.SetConfigType("json")
	// vp.AddConfigPath(".")

	// err := vp.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(vp.GetString("foo"))

	// vp.Set("name", "Marko")

	// vp.WriteConfig()

	// vp.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("File changed", in.Name)
	// })

	// vp.WatchConfig()

	// for {
	// }

	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(config)

}
