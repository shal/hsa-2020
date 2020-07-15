package config

import "fmt"

// Cache represents settings for connection to the cache.
type Cache struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}

// Addr returns cache address in {host}:{port} format.
func (r *Cache) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
