package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/shal/hsa-2020/04/pkg/cache/rediscache"

	"github.com/shal/hsa-2020/04/pkg/apiserver"
	"github.com/shal/hsa-2020/04/pkg/config"
	"github.com/shal/hsa-2020/04/pkg/store/mongostore"
)

func main() {
	var port int
	var configPath string

	flag.IntVar(&port, "port", 8080, "Port of the server")
	flag.StringVar(&configPath, "config", "config/config.toml", "Path to the configuration file")

	flag.Parse()

	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	store, err := mongostore.New(context.Background(), conf.Store)
	if err != nil {
		log.Fatal(err)
	}

	var cache *rediscache.Cache
	if conf.Cache.Ebabled {
		log.Println("Caching enabled")

		cache, err = rediscache.New(context.Background(), conf.Cache)
		if err != nil {
			log.Fatal(err)
		}
	}

	srv := apiserver.New(store, cache, false)
	addr := ":" + strconv.Itoa(port)
	log.Printf("Listening on %s...", addr)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}
