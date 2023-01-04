/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"my-web/internal/conf"
	"my-web/internal/log"
	"my-web/internal/server"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-web",
	Short: "一个cmd-web服务器例子",
	Long:  `长描述-cmd-web服务器例子`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("hello world")
		err := manageAPI()
		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("I'm inside init function in cmd/root.go")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&conf.ConfigFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&conf.WorkDir, "work-dir", "p", ".", "current work directory")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(
		newVersionCommand(),
	)
}

func manageAPI() error {
	conf.InitConf()
	//配置都初始化到 conf.go 中全局变量
	fmt.Printf("server port: %v \n", conf.ServerPort)
	log.InitLogger()

	s, err := server.NewServer(&server.Options{})
	if err != nil {
		return err
	}

	// start my-web API server
	errSig := make(chan error, 5)
	s.Start(errSig)

	// Signal received to the process externally.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Infof("The my-web server receive %s and start shutting down", sig.String())
		s.Stop()
		log.Infof("See you next time!")
	case err := <-errSig:
		log.Error("The my-web server start failed: %s", err.Error())
	}
	return nil
}
