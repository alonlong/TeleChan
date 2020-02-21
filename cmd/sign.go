package cmd

import (
	"TeleChan/pkg/api"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	// command parameters for SignUp subcommand
	SignUpCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	SignUpCmd.PersistentFlags().StringVar(&email, "email", "alonlong@163.com", "User's email for SignUp")
	SignUpCmd.PersistentFlags().StringVar(&password, "password", "123456", "User's passowrd for SignUp")

	// command parameters for SignIn subcommand
	SignInCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	SignInCmd.PersistentFlags().StringVar(&email, "email", "alonlong@163.com", "User's email for SignIn")
	SignInCmd.PersistentFlags().StringVar(&password, "password", "123456", "User's passowrd for SignIn")

	// command parameters for Refresh subcommand
	RefreshCmd.PersistentFlags().StringVar(&address, "address", "localhost:15001", "TeleChan server's address")
	RefreshCmd.PersistentFlags().StringVar(&token, "token", "", "User's token for Refresh")
}

func init() {
	RootCmd.AddCommand(SignUpCmd)
	RootCmd.AddCommand(SignInCmd)
	RootCmd.AddCommand(RefreshCmd)
}

// new a grpc client with address
func newClient(address string) (api.TeleChanClient, error) {
	// create the grpc connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// create the client with connection
	return api.NewTeleChanClient(conn), nil
}

// SignUpCmd subcommand
var SignUpCmd = &cobra.Command{
	Use:   "SignUp",
	Short: "User signs up the TeleChan",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.SignUpRequest{
			User: &api.User{
				Email:    email,
				Password: password,
			},
		}
		// call the SignUp interface
		if _, err := client.SignUp(context.Background(), request); err != nil {
			fmt.Printf("Sign up: %v\n", err)
			return
		}
		fmt.Println("Sign up success")
	},
}

// sign in with email and password
func signin(client api.TeleChanClient) (string, error) {
	request := &api.SignInRequest{
		User: &api.User{
			Email:    email,
			Password: password,
		},
	}
	// call the SignIn interface
	reply, err := client.SignIn(context.Background(), request)
	if err != nil {
		return "", err
	}
	return reply.Token, nil
}

// SignInCmd subcommand
var SignInCmd = &cobra.Command{
	Use:   "SignIn",
	Short: "User signs in the TeleChan",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		token, err := signin(client)
		if err != nil {
			fmt.Printf("Sign in: %v\n", err)
			return
		}
		fmt.Printf("Got token: %v\n", token)
	},
}

// RefreshCmd subcommand
var RefreshCmd = &cobra.Command{
	Use:   "Refresh",
	Short: "User refreshes token from the TeleChan",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.RefreshRequest{}
		ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(authorization, token))
		// call the Refresh interface
		reply, err := client.Refresh(ctx, request)
		if err != nil {
			fmt.Printf("Refresh: %v\n", err)
			return
		}

		fmt.Printf("Got refresh token: %v\n", reply.RefreshToken)
	},
}
