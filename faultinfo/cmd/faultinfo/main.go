package main

import (
	"os"

	"github.com/charakoba-com/everything/faultinfo"
	flags "github.com/jessevdk/go-flags"
	"github.com/web-apps-tech/identity/log"
)

type options struct {
	Listen string `short:"l" long:"listen" default:":8080" description:"server listen address"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Error.Print(err)
		return 1
	}
	s := faultinfo.New()
	log.Info.Printf("server listening at %s", opts.Listen)
	if err := s.Run(opts.Listen); err != nil {
		log.Error.Print(err)
		return 1
	}
	return 0
}
