/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jatalocks/opsilon/pkg/web"
	"github.com/spf13/cobra"
)

var port int64

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs an api server that functions the same as the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		web.App(port, ver)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serverCmd.Flags().Int64VarP(&port, "port", "p", 8080, "Port to start the web server in")
	// viper.BindPFlag("kubernetes", serverCmd.Flags().Lookup("kubernetes"))
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
