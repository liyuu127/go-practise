package cmd

import (
	"errors"
	"fmt"
	"github.com/liyuu127/go-practise/cobra/newApp/log"
	"github.com/spf13/cobra"
)

var hello = &cobra.Command{
	Use:   "hello",
	Short: "inner valid arg func",
	// Args:  cobra.MinimumNArgs(1),
	Args: func(cmd *cobra.Command, args []string) error { // 自定义验证函数
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Sugar.Infof("hello word %v\n", args[0])
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(hello)
}
