package main

import (
	"github.com/fsnotify/fsnotify"
	. "github.com/liyuu127/go-practise/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	cfg := pflag.StringP("config", "c", "", "Configuration file.")
	help := pflag.BoolP("help", "h", false, "Show this help message.")
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	if *cfg != "" {
		Sugar.Infof("cfg:%v", *cfg)
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("json")
	} else {
		// Viper 会根据添加的路径顺序搜索配置文件，如果找到则停止搜索
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
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
	Sugar.Infof("datastore.metric.host :%v", viper.GetString("datastore.metric.host"))

	type config struct {
		Port    int               `mapstructure:"host.port"`
		address string            `mapstructure:"host.address"`
		PathMap map[string]string `mapstructure:"host"`
	}
	var c config
	if err := viper.Unmarshal(&c); err != nil {
		Sugar.Errorf("Unmarshal error:%v", err)
	}

	Sugar.Infof("c:%+v", c)
	Sugar.Infof("config:\n%v", yamlStringSettings())

}

func yamlStringSettings() string {
	c := viper.AllSettings()
	bytes, err := yaml.Marshal(c)
	if err != nil {
		Sugar.Errorf("Unmarshal error:%v", err)
	}
	return string(bytes)

}
