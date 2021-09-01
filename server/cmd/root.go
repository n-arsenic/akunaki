package cmd

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/n-arsenic/grpc-service/server/lib/grpcserver"
	"github.com/n-arsenic/grpc-service/server/lib/system"
	"github.com/spf13/cobra"
)

type Config struct {
	Port    int
	Message string
}

func NewRootCmd() *cobra.Command {
	var config Config

	cmd := &cobra.Command{
		Use:   "judas",
		Short: "grpc server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if config.Message == "" {
				return errors.New("--message must be set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRootCmd(config)
		},
	}

	cmd.Flags().StringVarP(&config.Message, "message", "m", "", "server hello message")
	cmd.Flags().IntVarP(&config.Port, "port", "p", 8080, "server port")

	return cmd
}

func runRootCmd(config Config) error {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpcserver.NewGRPCServer(config.Message)
	system.Interrupt(server.GracefulStop)
	log.Printf("server listening at %v", conn.Addr())
	if err := server.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
