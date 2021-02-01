package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"portdb.io/src/proto/grpc"
)

type client struct {
	cli grpc.PortDomainServiceClient
}

func (c *client) Store(req *grpc.CreateRequest) (*grpc.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.cli.Store(ctx, req)
}

func (c *client) Fetch(req *grpc.FetchRequest) (*grpc.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.cli.Fetch(ctx, req)
}

func (c *client) Decoder(reader io.Reader) error {
	buf := bufio.NewReader(reader)
	dec := json.NewDecoder(buf)

	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v\n", t, t)
		var msg grpc.CreateRequest
		err = dec.Decode(&msg)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			return nil
		}
		if _, err := c.Store(&msg); err != nil {
			return err
		}
	}

	return nil
}

func New(cli grpc.PortDomainServiceClient) *client {
	return &client{cli: cli}
}
