package redis

import (
	"time"

	"github.com/tumiao/redis/internal/pool"
)

func (c *baseClient) Pool() pool.Pooler {
	return c.connPool
}

func (c *PubSub) Pool() pool.Pooler {
	return c.base.connPool
}

func (c *PubSub) ReceiveMessageTimeout(timeout time.Duration) (*Message, error) {
	return c.receiveMessage(timeout)
}
