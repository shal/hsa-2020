package config

import (
	"fmt"
	"strings"
)

type StoreNode struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// Store represents settings for connection to the MongoDB.
type Store struct {
	User       string      `toml:"username"`
	Password   string      `toml:"password"`
	Database   string      `toml:"database"`
	ReplicaSet *string     `toml:"replica_set"`
	Nodes      []StoreNode `toml:"nodes"`
}

// URI returns prepared URI to the mongodb.
func (m *Store) URI() string {
	nodes := make([]string, len(m.Nodes))
	for i, node := range m.Nodes {
		nodes[i] = fmt.Sprintf("%s:%d", node.Host, node.Port)
	}

	part := strings.Join(nodes, ",")
	addr := fmt.Sprintf("mongodb://%s:%s@%s/", m.User, m.Password, part)

	if m.ReplicaSet != nil && *m.ReplicaSet != "" {
		addr += "?replicaSet=" + *m.ReplicaSet
	}

	return addr
}
