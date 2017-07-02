package gate

import (
	"sync"

	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/constants"
	"github.com/hugohuang1111/treehole/db"
	"github.com/hugohuang1111/treehole/module"
	"github.com/hugohuang1111/treehole/router"
	"github.com/hugohuang1111/treehole/utils"
)

var (
	connIDMutex   sync.Mutex
	connIDCounter uint64
)

type connect struct {
	id       uint64
	closeSig chan bool
	sendChan chan []byte
	recvChan chan map[string]interface{}
}

func newConn() *connect {
	c := new(connect)

	connIDMutex.Lock()
	connIDCounter++
	c.id = connIDCounter
	connIDMutex.Unlock()
	c.closeSig = make(chan bool, 1)
	c.sendChan = make(chan []byte, 8)
	c.recvChan = make(chan map[string]interface{}, 8)
	go c.work()

	return c
}

func (c *connect) ID() uint64 {
	return c.id
}

func (c *connect) CloseSig() chan bool {
	return c.closeSig
}

func (c *connect) SendChan() chan []byte {
	return c.sendChan
}

func (c *connect) RecvChan() chan map[string]interface{} {
	return c.recvChan
}

func (c *connect) work() {
	for {
		msg := <-c.recvChan

		m := new(module.Mail)
		m.Sender = constants.ModGate
		m.Type = module.MailTypeNormal
		m.Payload = make(map[string]interface{})
		m.Payload[module.PayloadKeyConnectID] = c.id
		cmd := utils.GetStringFromMap(msg, "cmd")
		switch cmd {
		case "express":
			m.Recver = constants.ModDB
			m.Payload["cmd"] = db.DBCmdSavedWord
			m.Payload["word"] = utils.GetStringFromMap(msg, "word")
			m.Payload["nickName"] = utils.GetStringFromMap(msg, "nickName")
		case "topexpress":
			m.Recver = constants.ModDB
			m.Payload["cmd"] = db.DBCmdTopWords
			m.Payload["start"] = utils.GetIntFromMap(msg, "start")
			m.Payload["length"] = utils.GetIntFromMap(msg, "length")
			glog.Info(msg)
			glog.Info(m.Payload)
		default:
			glog.Info("Gate unhandler connect cmd ", cmd)
			continue
		}

		router.Route(m)
	}
}
