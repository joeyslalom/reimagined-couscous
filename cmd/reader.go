/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

// readerCmd represents the reader command
var readerCmd = &cobra.Command{
	Use:   "reader",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reader called")
		doReader()
	},
}

func init() {
	rootCmd.AddCommand(readerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doReader() {
	fmt.Println("Hello, world.")

	ctx := context.Background()
	projectId := "slalom-2020-293920"
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	subId := "reimagined-couscous-sub"
	sub := client.Subscription(subId)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		time.Sleep(10 * time.Second)
		m.Ack()
	})
	if err != nil {
		log.Fatalf("sub.Receive: %v", err)

	}
	cancel()
	log.Println("fin")
}