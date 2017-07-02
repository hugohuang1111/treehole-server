package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/treehold"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Info("start ...")

	treehold.Run()

	glog.Info("end ...")
}
