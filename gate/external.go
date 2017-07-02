package gate

import (
	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/gate/internal"
	"github.com/hugohuang1111/treehole/module"
)

//Gate gate module
type Gate struct {
	skelection module.Skelecton
}

//NewGate new gate module
func NewGate() module.Module {
	m := new(Gate)

	return m
}

//OnInit module init
func (g *Gate) OnInit() {
	g.skelection.OnInit(g)
	go internal.StartHTTPServer(
		internal.NewConnectHandler(newConnect),
		internal.DisConnectHandler(disConnect))
}

//OnDestroy module destroy
func (g *Gate) OnDestroy() {
	internal.StopHTTPServer()
	g.skelection.OnDestroy()
}

//OnMail module recv mail
func (g *Gate) OnMail(mail *module.Mail) {
	g.skelection.OnMail(mail)
}

//OnProcess process event
func (g *Gate) OnProcess(mail *module.Mail) {
	connID := module.GetConnectID(mail)
	switch mail.Type {
	case module.MailTypeSend:
		{
			if clientData := module.GetSendData(mail); nil != clientData {
				g.sendToClient(connID, clientData)
			}
		}
	default:
		glog.Warning("Gate unknow mail type:", mail.Type)
	}
}

func (g *Gate) sendToClient(connID uint64, data interface{}) {
	if sendData, ok := data.([]byte); ok {
		handSendData(connID, sendData)
	} else {
		glog.Warning("Gate sendToClient trans data failed")
	}
}
