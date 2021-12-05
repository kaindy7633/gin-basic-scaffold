package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量，保存程序的所有配置信息
var Conf = new(AppConfig)

// 定义 struct 映射配置信息
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	// 配置文件的两种方式
	// 方式1： 直接指定配置文件路径（相对路径或绝对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	// 绝对路径：系统中实际的文件路径
	viper.SetConfigFile(filePath)

	// 方式2：指定配置文件名和配置文件的位置，viper 自行查找可用的配置文件
	// 配置文件名不需要带后缀
	// 配置文件位置可配置多个

	// viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	// viper.SetConfigType("yaml")   // 指定配置文件类型
	// viper.AddConfigPath(".")      // 指定查找配置文件的路径

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal failed, err: #{err}\n")
	}

	// 热加载配置修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config modified ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper Unmarshal failed, err: #{err}\n")
		}
	})

	return
}
