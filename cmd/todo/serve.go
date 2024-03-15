package main

import (
	"context"
	"errors"
	"fmt"
	pb "git.local/jmercado/todo/api"
	"git.local/jmercado/todo/server"
	"github.com/oklog/run"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"os"
	"syscall"
	"time"
)

type options struct {
	gRPC    bool
	devMode bool
}

func commandServe() *cobra.Command {

	options := options{devMode: false, gRPC: true}

	cmd := &cobra.Command{
		Use:     "serve",
		Short:   "Start Todo Server",
		Example: "todo serve",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true

			err := runServe(options)
			return err
		},
	}

	return cmd
}

func runServe(options options) error {

	var group run.Group

	appServer, err := server.NewServer(context.Background(), options.devMode)

	if err != nil {
		return fmt.Errorf("error start server: %v", err)
	}

	addr := "127.0.0.1:8080"
	log.Printf("http server listening on: %s", addr)

	group.Add(func() error {
		return appServer.Start(context.Background(), addr)
	}, func(err error) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		log.Info().Msg("Starting server shutdown")
		_ = appServer.Stop(ctx)
	})

	if options.gRPC {
		apiServer, err := server.NewApiServer(context.Background())

		if err != nil {
			return fmt.Errorf("error creating gRPC server: %v", err)
		}

		addr := "127.0.0.1:5104"
		grpcListener, err := net.Listen("tcp", addr)

		if err != nil {
			return fmt.Errorf("error creating gRPC listener: %v", err)
		}

		defer grpcListener.Close()
		log.Info().Msgf("gRPC listening on: %s", addr)

		grpcServer := grpc.NewServer()
		pb.RegisterTodoServiceServer(grpcServer, apiServer)

		group.Add(
			func() error {
				return grpcServer.Serve(grpcListener)
			},
			func(err error) {
				_, cancel := context.WithTimeout(context.Background(), time.Second*20)
				defer cancel()

				log.Info().Msg("Starting shutdown gRPC server")
				grpcServer.GracefulStop()
			})
	}

	group.Add(run.SignalHandler(context.Background(), os.Interrupt, syscall.SIGTERM))

	if err := group.Run(); err != nil {
		var signalError run.SignalError

		if !errors.As(err, &signalError) {
			return fmt.Errorf("error in run groups: %w", err)
		}

		log.Printf("%v, shutdown now", err)
	}

	return nil
}
