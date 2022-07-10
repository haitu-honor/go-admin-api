package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/myadmin/project/global"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 { // path是一个string切片
		flag.StringVar(&config, "c", "", "choose config file.") // 选择config配置文件
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" { // 判断 ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() { // mode 可以通过 gin.setMode() 进行设置，默认是debug
				case gin.DebugMode:
					config = ConfigDefaultFile
					fmt.Printf("您正在使用gin.DebugMode环境名称,config的路径为%s\n", ConfigDefaultFile)
				case gin.ReleaseMode:
					config = ConfigReleaseFile
					fmt.Printf("您正在使用gin.ReleaseMode环境名称,config的路径为%s\n", ConfigReleaseFile)
				case gin.TestMode:
					config = ConfigTestFile
					fmt.Printf("您正在使用gin.TestMode环境名称,config的路径为%s\n", ConfigTestFile)
				}
			} else { // ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config) // 参数config是配置文件路径, "./config.yml" ./ 代表根目录下
	v.SetConfigType("yaml") // 设置配置文件类型
	err := v.ReadInConfig() // 加载yml配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig() // 动态监测配置文件
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生改变---:", e.Name)
		if err = v.Unmarshal(&global.GAI_CONFIG); err != nil {
			fmt.Println("配置重载失败---:", err)
		}
	})
	// 将config.yml配置信息解析到 global.GAI_CONFIG 结构体
	if err = v.Unmarshal(&global.GAI_CONFIG); err != nil {
		fmt.Println("配置重载失败---:", err)
	}

	return v
}
