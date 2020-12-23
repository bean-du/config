package main

import (
	"coco-tool/config/conf"
	"coco-tool/config/repostory"
	"coco-tool/config/web"
	"errors"
	"os"
	"os/signal"
	"syscall"

	_ "coco-tool/config/provider"
	_ "coco-tool/config/provider/etcd"
	"github.com/urfave/cli/v2"
)

var cancelFuncs = make([]func(), 0, 0)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"f"},
				Usage:   "application config file",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("config") == "" {
				return errors.New("config is empty")
			}
			cancelFuncs = append(cancelFuncs, conf.Init(c.String("config"))...)
			cancelFuncs = append(cancelFuncs, web.Run())
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
	repostory.Init()
	defer func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
