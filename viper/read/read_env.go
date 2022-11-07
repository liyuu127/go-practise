package main

import (
	. "github.com/liyuu127/go-practise/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func main() {
	port := pflag.String("port", "", "help for java port")
	pflag.Parse()
	// viper.SetEnvPrefix("VIPER")
	viper.BindEnv("home", "home")
	viper.BindEnv("username", "username")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	Sugar.Infof("home:%v", viper.Get("home"))
	Sugar.Infof("javaHome:%v", viper.Get("java.home"))
	Sugar.Infof("javaHome:%v", viper.Get("java_home"))
	Sugar.Infof("javaHome:%v", viper.Get("JAVA_HOME"))

	// 使用环境变量
	os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
	os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb")

	// viper.AutomaticEnv()                                             // 读取环境变量
	viper.SetEnvPrefix("VIPER") // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写。
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
	viper.BindEnv("user.secret-key")
	viper.BindEnv("user.secret-id", "USER_SECRET_ID") // 绑定环境变量名到key
	Sugar.Infof("user.secret:%v", viper.Get("user.secret-key"))
	Sugar.Infof("user.secret-id:%v", viper.Get("user.secret-id"))

	Sugar.Infof("port:%v", *port)
	// viper.BindPFlag("port", pflag.Lookup("port")) // 绑定单个标志
	viper.BindPFlags(pflag.CommandLine) // 绑定标志集
	Sugar.Infof("port:%v", viper.Get("port"))

}
