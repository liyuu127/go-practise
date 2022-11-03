package main

import (
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"strings"
)

var (
	flagVar  = pflag.IntP("flagName", "f", 1234, "help message for flagname")
	name     = pflag.StringP("name", "a", "liyu", "help message for name")
	password = pflag.StringP("password", "p", "123456", "help message for password")
	port     = pflag.IntP("port", "P", 3306, "MySQL service host port.")
	help     = pflag.BoolP("help", "h", false, "help message")
	ggh      = pflag.BoolP("gg-h", "g", false, "help message")
)

func main() {
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalize)
	// NoOptDefVal
	pflag.Lookup("flagName").NoOptDefVal = "4321"

	// Deprecated
	pflag.CommandLine.MarkDeprecated("name", "use --flagName instead")

	// short deprecated
	pflag.CommandLine.MarkShorthandDeprecated("port", "usr --port only")

	// hidden
	pflag.CommandLine.MarkHidden("password")

	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	logger, _ := zap.NewDevelopment(zap.AddCaller())
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Infof("flagVar=%d", *flagVar)

	flagSet := pflag.CommandLine
	flagName, err := flagSet.GetInt("flagName")
	if err != nil {
		sugar.Error(err)
	}
	sugar.Infow("get flagName",
		"flagName", flagName,
	)

	// no arg option
	sugar.Infof("arg number is :%v", pflag.NArg())
	sugar.Infof("arg list is :%v", pflag.Args())
	sugar.Infof("arg first is :%v", pflag.Arg(0))

	_ggh, err := flagSet.GetBool("gg.h")
	sugar.Infof("ggh :%v", _ggh)
	sugar.Infof("ggh :%v", *ggh)

}

func wordSepNormalize(set *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}
