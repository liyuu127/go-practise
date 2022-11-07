package cmd

import (
	. "github.com/liyuu127/go-practise/cobra/newApp/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile     string
	projectBase string
	userLicense string
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with love by spf13 and friends in Go. 
				Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Sugar.Fatalf("execute error:%-v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// 标志可以是“持久的”，这意味着该标志可用于它所分配的命令以及该命令下的每个子命令
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// 本地标志只能在它所绑定的命令上使用
	rootCmd.Flags().StringP("source", "s", "", "Source directory to read from")

	// 设置标志为必选
	// rootCmd.MarkFlagRequired("region")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE ")
	viper.SetDefault("license", "apache")

	rootCmd.SuggestionsMinimumDistance = 2
}

func initConfig() {
	if cfgFile != "" {
		Sugar.Infof("set config:%v", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		// dir, err := homedir.Dir()
		// if err != nil {
		// 	Sugar.Fatalf("get dir error:%-v", err)
		// 	os.Exit(1)
		// }

		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		Sugar.Fatalf("can not read config error:%-v", err)
		os.Exit(1)
	}
}
