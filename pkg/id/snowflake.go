package id

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

func InitSnowflake(nodeID int64) error {

	var err error

	once.Do(func() {
		node, err = snowflake.NewNode(nodeID)
	})

	return err
}

func GenerateID() int64 {

	if node == nil {
		panic("snowflake not initialized")
	}

	return node.Generate().Int64()
}
