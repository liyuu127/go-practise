package main

import (
	"github.com/fsnotify/fsnotify"
	. "github.com/liyuu127/go-practise/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	if *cfg != "" {
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("yaml")
	} else {
		// Viper 会根据添加的路径顺序搜索配置文件，如果找到则停止搜索
		viper.AddConfigPath("../")
		viper.SetConfigName("outer_config")
	}
	if err := viper.ReadInConfig(); err != nil {
		Sugar.Error(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		Sugar.Infof("Config file changed:%v", e.Name)
	})

	Sugar.Infof("used config :%v", viper.ConfigFileUsed())

}
