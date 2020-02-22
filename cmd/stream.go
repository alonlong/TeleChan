package cmd

import (
	"TeleChan/pkg/api"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

func init() {
	// command parameters for Communicate subcommand
	CommunicateCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	CommunicateCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for Communicate")

	// command parameters for Publish subcommand
	PublishCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	PublishCmd.PersistentFlags().StringVar(&channelName, "name", "Go", "The channel name for Publish")
	PublishCmd.PersistentFlags().StringVar(&content, "content", "This is a new message", "The message content for Publish")
	PublishCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for Publish")
}

func init() {
	RootCmd.AddCommand(CommunicateCmd)
	RootCmd.AddCommand(PublishCmd)
}

// PublishCmd subcommand
var PublishCmd = &cobra.Command{
	Use:   "Publish",
	Short: "User publish a message for channel",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.PublishRequest{
			Name:    channelName,
			Content: content,
		}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// call the Publish interface
		if _, err := client.Publish(ctx, request); err != nil {
			fmt.Printf("Publish message: %v\n", err)
			return
		}
		fmt.Println("Publish message success")
	},
}

// CommunicateCmd subcommand
var CommunicateCmd = &cobra.Command{
	Use:   "Communicate",
	Short: "User communicate with TeleChan server",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// new stream from client
		stream, err := client.Communicate(ctx)
		if err != nil {
			fmt.Printf("New stream: %v\n", err)
			return
		}

		connect := &api.ClientUpdate{
			Message: &api.ClientUpdate_Up{
				Up: &api.Connect{},
			},
		}
		// send the connect message first
		if err := stream.Send(connect); err != nil {
			fmt.Printf("Send connect message: %v\n", err)
			return
		}

		// start a goroutine to recv the pushing messages
		go func() {
			for {
				reply, err := stream.Recv()
				if err != nil {
					fmt.Printf("Stream recv: %v\n", err)
					os.Exit(-1)
				}

				data := reply.GetData()
				if data != nil {
					fmt.Printf("Receive message: %v\n", data)
				}

				pong := reply.GetPong()
				if pong != nil {
					latency := time.Now().UnixNano()/1000000 - int64(pong.Timestamp)
					fmt.Printf("Ping round trip, lantency: %dms\n", latency)
				}
			}
		}()

		// start a goroutine to send ping message
		go func() {
			for {
				select {
				case <-time.After(time.Second * 5):
					ping := &api.ClientUpdate{
						Message: &api.ClientUpdate_Ping{
							Ping: &api.Ping{
								Timestamp: uint64(time.Now().UnixNano() / 1000000),
							},
						},
					}
					if err := stream.Send(ping); err != nil {
						fmt.Printf("Send ping: %v\n", err)
						os.Exit(-1)
					}
				}
			}
		}()

		select {
		case <-time.After(60 * time.Minute):
			stream.CloseSend()

			fmt.Println("shut down the stream")
			return
		}
	},
}
