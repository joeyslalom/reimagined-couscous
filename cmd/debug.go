/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("debug called")
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			log.Fatalln("debug.ReadBuildInfo() failed")
		}
		log.Printf("buildInfo path:%v main.path:%v main.version:%v main.sum:%v go:%v", bi.Path, bi.Main.Path, bi.Main.Version, bi.Main.Sum, bi.GoVersion)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debugCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
