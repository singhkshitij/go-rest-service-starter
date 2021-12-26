package cmd

import (
	"fmt"
	"github.com/singhkshitij/golang-rest-service-starter/src/cache"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "golang-rest-service-starter",
	Short: "A RESTful API boilerplate",
	Long:  `A RESTful API boilerplate for golang`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initSetup)
	cfgFile := config.GetConfigFile()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config.yaml", "", "config file (default is $HOME/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initSetup() {
	config.InitConfig()
	logger.Must(logger.NewLogger("")) //replace with config file if logs need to be put in a log file
	err := cache.Setup()
	if err != nil {
		panic("err while setting up redis " + err.Error())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
