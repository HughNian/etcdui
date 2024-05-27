package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HughNian/nmid-discovery/conf"
	"github.com/HughNian/nmid-discovery/e3ch"
	"github.com/HughNian/nmid-discovery/routers"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/api/v3/version"
)

const (
	PROGRAM_NAME    = "etcdui"
	PROGRAM_VERSION = "0.1.0"
)

var configFilepath string

func init() {
	flag.StringVar(&configFilepath, "conf", "conf/config.default.ini", "config file path")
	rev := flag.Bool("rev", false, "print rev")
	flag.Parse()

	if *rev {
		fmt.Printf("[%s v%s]\n[etcd %s]\n",
			PROGRAM_NAME, PROGRAM_VERSION,
			version.Version,
		)
		os.Exit(0)
	}
}

func main() {
	config, err := conf.Init(configFilepath)
	if err != nil {
		panic(err)
	}

	client, err := e3ch.NewE3chClient(config)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.UseRawPath = true
	routers.InitRouters(router, config, client)
	if err := router.Run(":" + config.Port); err != nil {
		panic(err)
	}
}
