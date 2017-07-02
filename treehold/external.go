package treehold

import (
	"github.com/hugohuang1111/treehole/constants"
	"github.com/hugohuang1111/treehole/db"
	"github.com/hugohuang1111/treehole/gate"
	"github.com/hugohuang1111/treehole/module"
	"github.com/hugohuang1111/treehole/router"
)

// Run run
func Run() {
	close := make(chan bool, 1)

	router.Start()
	module.Register(constants.ModGate, gate.NewGate())
	module.Register(constants.ModDB, db.NewDB())
	module.Run()

	<-close
	router.End()
}
