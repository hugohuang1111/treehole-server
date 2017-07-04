package gate

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/constants"
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
	switch mail.Type {
	case module.MailTypeSendToClient:
		connID := module.GetConnectID(mail)
		cmd := module.GetPayloadValueString(mail, module.PayloadKeyCmd)
		sendMap := make(map[string]interface{})
		switch cmd {
		case constants.CmdDBTopWords:
			sendMap["words"] = module.GetPayloadValueStringArr(mail, "words")
			fallthrough
		case constants.CmdDBSaveWord:
			e := module.GetPayloadValueInt(mail, module.PayloadKeyError)
			sendMap["error"] = e
			sendMap["description"] = constants.ErrorDescription[e]
			sendMap["words"] = module.GetPayloadValue(mail, "words")
		default:
			sendMap["error"] = constants.ErrFailed
			sendMap["description"] = constants.ErrorDescription[constants.ErrFailed]
			glog.Warning("Gate unknow cmd:", cmd)
		}

		if bytes, err := json.Marshal(sendMap); nil == err {
			g.sendToClient(connID, bytes)
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
