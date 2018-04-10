package main

import (
	"log"
	"os"

	"github.com/charakoba-com/everything/faultinfo"
	"github.com/charakoba-com/everything/faultinfo/db"
	"github.com/go-sql-driver/mysql"
	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Listen string `short:"l" long:"listen" default:":8080" description:"server listen address"`
	MySQL  mysqlOptions
}

type mysqlOptions struct {
	Addr     string `long:"mysql-addr" default:"127.0.0.1:3306"`
	User     string `long:"mysql-user" default:"root"`
	Password string `long:"mysql-password" default:""`
	DBName   string `long:"mysql-db" default:"faultinfo"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Print(err)
		return 1
	}
	mysqlCfg := mysql.Config{
		Net:       "tcp",
		Addr:      opts.MySQL.Addr,
		User:      opts.MySQL.User,
		Passwd:    opts.MySQL.Password,
		DBName:    opts.MySQL.DBName,
		ParseTime: true,
	}
	if err := db.Open(mysqlCfg.FormatDSN()); err != nil {
		log.Print(err)
		return 1
	}
	defer db.Close()
	s := faultinfo.New(opts.Listen)
	if err := s.Run(); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}
