package cmd

import (
	"TeleChan/pkg/api"
	"TeleChan/pkg/config"
	"TeleChan/pkg/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

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
	service := api.NewService(&settings)

	var wg sync.WaitGroup
	// start a server with listen address
	service.Start(&wg)

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

			// stop the grpc server gracefully
			service.Stop()

			log.Info("server is stopped")

			// wait for server goroutine exit first
			wg.Wait()

			log.Infof("shut down takes time: %v", time.Now().Sub(start))
			return
		}
	}
}
