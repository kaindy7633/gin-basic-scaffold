package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")   // 指定配置文件类型
	viper.AddConfigPath(".")      // 指定查找配置文件的路径
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}

	// 热加载配置修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config modified ...")
	})

	return
}