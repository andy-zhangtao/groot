package main

import (
	"github.com/andy-zhangtao/bwidow"
	"github.com/andy-zhangtao/gogather/zlog"
	"github.com/sirupsen/logrus"
)

var bw *bwidow.BW
var z *zlog.Zlog

const ModuleName = "Groot"

func init() {
	z = zlog.GetZlog()
	bw = bwidow.GetWidow()
	if err := bw.Driver(bwidow.DRIVER_PQ).Error(); err != nil {
		panic(err)
	}

	if err := bw.Map(Groot{}, GROOT_ACCOUNT_DB).Error(); err != nil {
		panic(err)
	}

	if err := bw.CheckIndex(new(Groot)).Error(); err != nil {
		panic(err)
	}

	logrus.WithFields(z.Fields(logrus.Fields{"BW Init Success Version": bw.Version()})).Info(ModuleName)
}
