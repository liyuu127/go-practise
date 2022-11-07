package main

import (
	. "github.com/liyuu127/go-practise/viper"
	"github.com/spf13/viper"
)

// 通过 viper.Set 函数显示设置的配置
// 命令行参数
// 环境变量
// 配置文件
// Key/Value
// 存储默认值
func main() {

	// 设置默认值。
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.Set("user.username", "colin")
	viper.Set("LayoutDir", "aaa")

	Sugar.Infow("viper read",
		"ContentDir", viper.Get("ContentDir"),
		"LayoutDir", viper.Get("LayoutDir"),
		"Taxonomies", viper.Get("Taxonomies"),
		"user.username", viper.Get("user.username"),
	)
}
