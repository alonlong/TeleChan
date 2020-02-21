package cmd

import (
	"TeleChan/pkg/api"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

func init() {
	// command parameters for NewChan subcommand
	NewChanCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	NewChanCmd.PersistentFlags().StringVar(&channelName, "name", "Go", "User's channel name for NewChan")
	NewChanCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for NewChan")

	// command parameters for JoinChan subcommand
	JoinChanCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	JoinChanCmd.PersistentFlags().StringVar(&channelName, "name", "Go", "User's channel name for JoinChan")
	JoinChanCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for JoinChan")

	// command parameters for GetChans subcommand
	GetChansCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	GetChansCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for GetChans")
}

func init() {
	RootCmd.AddCommand(NewChanCmd)
	RootCmd.AddCommand(JoinChanCmd)
	RootCmd.AddCommand(GetChansCmd)
}

// NewChanCmd subcommand
var NewChanCmd = &cobra.Command{
	Use:   "NewChan",
	Short: "User create channel",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.NewChanRequest{
			Name: channelName,
		}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// call the NewChan interface
		if _, err := client.NewChan(ctx, request); err != nil {
			fmt.Printf("New channel: %v\n", err)
			return
		}
		fmt.Println("New channel success")
	},
}

// JoinChanCmd subcommand
var JoinChanCmd = &cobra.Command{
	Use:   "JoinChan",
	Short: "User join a special channel",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.JoinChanRequest{
			Name: channelName,
		}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// call the JoinChan interface
		if _, err := client.JoinChan(ctx, request); err != nil {
			fmt.Printf("Join channel: %v\n", err)
			return
		}
		fmt.Println("Join channel success")
	},
}

// GetChansCmd subcommand
var GetChansCmd = &cobra.Command{
	Use:   "GetChans",
	Short: "User get channels",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.GetChansRequest{}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// call the GetChans interface
		reply, err := client.GetChans(ctx, request)
		if err != nil {
			fmt.Printf("Get channels: %v\n", err)
			return
		}
		for _, channel := range reply.Channels {
			fmt.Printf("Channel id: %s, name: %s\n", channel.Id, channel.Name)
		}
	},
}
