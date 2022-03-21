package cmd

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/joeyslalom/reimagined-couscous/proto"
)

// writerCmd represents the writer command
var writerCmd = &cobra.Command{
	Use:   "writer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("writerCmd.Run()")
		doWriter()
	},
}

func init() {
	rootCmd.AddCommand(writerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doWriter() {
	log.Println("doWriter()")

	ctx := context.Background()
	projectId := "slalom-2020-293920"
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	topicId := "reimagined-couscous"
	t := client.Topic(topicId)
	msg := &pb.PubsubPayload{
		Type:                 pb.PubsubPayload_IMAGE,
		Avatar:               &pb.PubsubPayload_ImageUrl{ImageUrl: "the-url"},
		Nested:               &pb.PubsubPayload_Nested{Name: "Nested Name", Id: 2021},
		Nums:                 []int32{10,20,30},
		NestedMap:            map[string]*pb.PubsubPayload_Nested{
			"uno": {Name: "one", Id: 123},
			"dos": {Name: "two", Id: 234},
		},
		Completed:            true,
		LastUpdated:          timestamppb.Now(),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("proto.Marshal: %v", err)
	}
	log.Printf("publishing message: %v", msg)
	log.Printf("timestamp: %v", msg.LastUpdated.AsTime())
	result := t.Publish(ctx, &pubsub.Message{Data: data})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("result.Get: %v", err)
	}
	log.Printf("Published a message id=%v", id)
}
