package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

const Version = "v0.1.26"

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login srun",
	RunE:  LoginE,
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout srun",
	RunE:  LogoutE,
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "get srun info",
	RunE:  InfoE,
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config srun",
	RunE:  ConfigE,
}

var rootCmd = &cobra.Command{
	Use:   "srun [command]",
	Short: "A efficient client for BIT campus network",
	RunE: func(cmd *cobra.Command, args []string) error {
		if debugMode {
			log.SetLevel(log.DebugLevel)
		}
		return LoginE(cmd, args)
	},
}

var debugMode bool

func main() {

	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "debug mode")

	rootCmd.Version = Version
	rootCmd.SetVersionTemplate(VersionString())

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(configCmd)

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		//DisableTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
