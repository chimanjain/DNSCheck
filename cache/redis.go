package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"os"
	"sync"

	"github.com/chimanjain/dnscheck/model"
	"github.com/redis/go-redis/v9"
)

var once sync.Once

type Client struct {
	client *redis.Client
}

var redisClient *Client

func GetRedisClient() *Client {
	if redisClient == nil {
		once.Do(
			func() {
				redisClient = NewRedis()
			})
	}
	return redisClient
}

func NewRedis() *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	return &Client{
		client: client,
	}
}

func (c *Client) GetDNS(ctx context.Context, url string) (model.Dns, error) {
	cmd := c.client.Get(ctx, url)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return model.Dns{}, err
	}

	b := bytes.NewReader(cmdb)

	var dns model.Dns

	if err := gob.NewDecoder(b).Decode(&dns); err != nil {
		return model.Dns{}, err
	}

	return dns, nil
}

func (c *Client) SetDNS(ctx context.Context, d model.Dns) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(d); err != nil {
		return err
	}

	return c.client.Set(ctx, d.URL, b.Bytes(), 0).Err()
}
