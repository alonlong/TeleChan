package cmd

import (
	"net"
	"os"
	"os/signal"
	"TeleChan/pkg/api"
	"TeleChan/pkg/config"
	"TeleChan/pkg/log"
	"sync"
	"syscall"
	"time"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// DefaultKeepaliveMinTime - if a client pings more than once every 5 seconds, terminate the connection
	DefaultKeepaliveMinTime = 5
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             DefaultKeepaliveMinTime * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,                                  // Allow pings even when there are no active streams
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		// start the grpc server
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

// start a grpc server
func startGRPCServer(settings *config.Config, impl *api.Service , wg *sync.WaitGroup) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", settings.Server.ListenAddr)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep))
	api.RegisterTeleChanServer(s, impl)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			log.Infof("grpc serve: %v\n", err)
		}
	}()

	return s, nil
}

// updateOptions updates the log options
func updateOptions(scope string, options *log.Options, settings *config.Config) error {
	if settings.Log.OutputPath != "" {
		options.OutputPaths = []string{settings.Log.OutputPath}
	}
	if settings.Log.RotationPath != "" {
		options.RotateOutputPath = settings.Log.RotationPath
	}

	options.RotationMaxBackups = settings.Log.RotationMaxBackups
	options.RotationMaxSize = settings.Log.RotationMaxSize
	options.RotationMaxAge = settings.Log.RotationMaxAge
	options.JSONEncoding = settings.Log.JSONEncoding
	level, err := options.ConvertLevel(settings.Log.OutputLevel)
	if err != nil {
		return err
	}
	options.SetOutputLevel(scope, level)
	options.SetLogCallers(scope, true)

	return nil
}

// the main process for the serve subcommand
func serve() {
	var settings config.Config
	// parse the config file
	if err := config.ParseYamlFile(configFile, &settings); err != nil {
		panic(err)
	}

	// init and update the log options
	logOptions := log.DefaultOptions()
	if err := updateOptions("default", logOptions, &settings); err != nil {
		panic(err)
	}
	// configure the log options
	if err := log.Configure(logOptions); err != nil {
		panic(err)
	}

	// create an api implemention for grpc server
	impl := api.NewService(&settings)

	var wg sync.WaitGroup
	// start a server with listen address
	grpcServer, err := startGRPCServer(&settings, impl, &wg)
	if err != nil {
		panic(err)
	}

	log.Info("server is started")

	sig := make(chan os.Signal, 1024)
	// subscribe signals: SIGINT & SINGTERM
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-sig:
			log.Infof("receive signal: %v", s)

			// flush the log
			log.Sync()

			start := time.Now()

			// close the grpc server gracefully
			grpcServer.GracefulStop()

			log.Info("server is stopped")

			// wait for server goroutine exit first
			wg.Wait()

			log.Infof("shut down takes time: %v", time.Now().Sub(start))
			return
		}
	}
}
