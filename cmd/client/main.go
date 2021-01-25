package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	"portdb.io/src/client"
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
		Path string
	}

	flag.StringVar(&cfg.Grpc.Host, "grpc-host", "server", "GRPC server host")
	flag.IntVar(&cfg.Grpc.Port, "grpc-port", 3000, "GRPC server port")
	flag.StringVar(&cfg.Path, "path", "/app/ports.json", "Path of the file that contains the ports")

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[1] != "" {
			switch pair[0] {
			case "GRPC-HOST":
				cfg.Grpc.Host = pair[1]
			case "GRPC-PORT":
				p, err := strconv.ParseInt(pair[1], 10, 32)
				if err != nil {
					return err
				}
				cfg.Grpc.Port = int(p)
			case "FILE-PATH":
				cfg.Path = pair[1]
			}
		}
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock())

	con, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer con.Close()
	cli := pb.NewPortDomainServiceClient(con)
	c := client.New(cli)
	fl, err := os.OpenFile(cfg.Path, os.O_RDONLY, 600)
	if err != nil {
		return fmt.Errorf("unable to read port file")
	}
	if err := c.Decoder(fl); err != nil {
		return err
	}
	return nil
}
