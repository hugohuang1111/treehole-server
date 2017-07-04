package gate

import (
	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/gate/internal"
)

var (
	connMap map[uint64]*connect
)

func init() {
	connMap = make(map[uint64]*connect)
}

func newConnect() internal.Connect {
	c := newConn()
	connMap[c.ID()] = c

	return c
}

func disConnect(connID uint64) {
	delete(connMap, connID)

	// m := new(module.Mail)
	// m.Recver = constants.ModGate
	// m.Sender = constants.ModGate
	// m.Type = module.MailTypeDisconnect
	// m.Payload = make(map[string]interface{})
	// m.Payload[module.PayloadKeyConnectID] = connID
	// router.Route(m)
}

func handSendData(connID uint64, payload []byte) {
	c, exist := connMap[connID]
	if !exist {
		glog.Warning("gate not find connect:", connID)
		return
	}
	ch := c.SendChan()
	ch <- payload
}
