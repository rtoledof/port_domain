package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"

	"portdb.io/src/port"
	"portdb.io/src/port/repo"
	pb "portdb.io/src/proto/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var cfg struct {
		Grpc struct {
			Host string
			Port int
		}
		DB struct {
			Host string
			Port int
			User string
			Pass string
			Name string
		}
	}
	flag.StringVar(&cfg.Grpc.Host, "grpc-host", "", "GRPC server host")
	flag.IntVar(&cfg.Grpc.Port, "grpc-port", 3000, "GRPC server port")
	flag.StringVar(&cfg.DB.Host, "db-address", "localhost", "Database server host")
	flag.IntVar(&cfg.DB.Port, "db-port", 3309, "Database server port")
	flag.StringVar(&cfg.DB.User, "db-user", "golang", "Database server user")
	flag.StringVar(&cfg.DB.Pass, "db-password", "golang", "Database server password")
	flag.StringVar(&cfg.DB.Name, "db-name", "port_domain", "Database server password")

	flag.Parse()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		switch pair[0] {
		case "GRPC-HOST":
			cfg.Grpc.Host = pair[1]
		case "GRPC-PORT":
			p, err := strconv.ParseInt(pair[1], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid grpc port address provided")
			}
			cfg.Grpc.Port = int(p)
		case "DB-HOST":
			cfg.DB.Host = pair[1]
		case "DB-PORT":
			p, err := strconv.ParseInt(pair[1], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid db port address provided")
			}
			cfg.DB.Port = int(p)
		case "DB-USER":
			cfg.DB.User = pair[1]
		case "DB-PASSWORD":
			cfg.DB.Pass = pair[1]
		case "DB-NAME":
			cfg.DB.Name = pair[1]
		}
	}

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name))
	if err != nil {
		return fmt.Errorf("cannot connect to %s: %v", cfg.DB.Name, err)
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGABRT)
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port))
	if err != nil {
		serverErrors <- err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPortDomainServiceServer(grpcServer, port.New(repo.NewMysql(db)))
	if err := grpcServer.Serve(lis); err != nil {
		serverErrors <- err
	}

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("stoping server: %v", err)

	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown", sig)

		// Asking listener to shutdown and load shed.
		grpcServer.Stop()

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return fmt.Errorf("integrity issue caused shutdown")
		}
	}

	return nil
}
