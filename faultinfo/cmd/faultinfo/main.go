package main

import (
	"log"
	"os"

	"github.com/charakoba-com/everything/faultinfo"
	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Listen string `short:"l" long:"listen" default:":8080" description:"server listen address"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Print(err)
		return 1
	}
	s := faultinfo.New(opts.Listen)
	if err := s.Run(); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}
