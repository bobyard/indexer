package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	APTOS = 0
	SUI   = 1
)

var Engine *xorm.Engine
var Sui *xorm.Engine

func ConnectSui(connStr string) {
	var err error
	Sui, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Panicf("%v", err)
	}

	f, err := os.Create("sui-db.log")
	if err != nil {
		println(err.Error())
		return
	}
	Sui.SetLogger(xorm.NewSimpleLogger(f))
}

func Connect(connStr string) {

	var err error
	Engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Panicf("%v", err)
	}
	f, err := os.Create("bobyard.log")
	if err != nil {
		println(err.Error())
		return
	}
	Engine.SetLogger(xorm.NewSimpleLogger(f))
}
