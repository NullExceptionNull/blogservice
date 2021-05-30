package order

import (
	opb "blog-service/api/order"
	"google.golang.org/grpc"
	"sync"
)

type Client struct {
	Addr   string
	once   sync.Once
	Client *opb.OrderServiceClient
}

func (c *Client) CreateClient() {
	c.once.Do(func() {
		conn, err := grpc.Dial(c.Addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		client := opb.NewOrderServiceClient(conn)
		defer conn.Close()
		c.Client = &client
	})

}
