package client

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"portdb.io/src/proto/grpc"
)

var re = regexp.MustCompile(`.*\"(?P<key>[A-Z]+)\"\:\s(?P<body>\{(.*\s*)((\"\w*\"(\:\s.*)*\s*)|((\-)*\d+(\.\d+).*\s*)|\].*\s*)+\})`)

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
	buf := make([]byte, 1024)
	var str string
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}
		str += string(buf[0:n])
		match := re.FindAllStringSubmatch(str, -1)
		if match != nil {
			for _, v := range match {
				var req grpc.CreateRequest
				dec := json.NewDecoder(strings.NewReader(v[2]))
				str = strings.Replace(str, v[2], "", 1)
				if err := dec.Decode(&req); err != nil {
					log.Println(err)
					continue
				}
				if _, err := c.Store(&req); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func New(cli grpc.PortDomainServiceClient) *client {
	return &client{cli: cli}
}
