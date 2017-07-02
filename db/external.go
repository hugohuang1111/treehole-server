package db

import (
	"github.com/golang/glog"
	"github.com/hugohuang1111/treehole/module"
)

const (
	//DBCmdSavedWord saved word
	DBCmdSavedWord = "saved word"

	//DBCmdTopWords top word
	DBCmdTopWords = "top words"
)

//DB db module
type DB struct {
	skelection module.Skelecton
}

//NewDB new gate module
func NewDB() module.Module {
	m := new(DB)

	return m
}

//OnInit module init
func (db *DB) OnInit() {
	db.skelection.OnInit(db)
}

//OnDestroy module destroy
func (db *DB) OnDestroy() {
	db.skelection.OnDestroy()
}

//OnMail module recv mail
func (db *DB) OnMail(mail *module.Mail) {
	db.skelection.OnMail(mail)
}

//OnProcess process event
func (db *DB) OnProcess(mail *module.Mail) {
	// connID := module.GetConnectID(mail)
	switch mail.Type {
	case module.MailTypeNormal:
		cmd := module.GetCmd(mail)
		switch cmd {
		case DBCmdSavedWord:
			word := module.GetPayloadValueString(mail, "word")
			nickName := module.GetPayloadValueString(mail, "nickName")
			saveword(nickName, word)
		case DBCmdTopWords:
			start := module.GetPayloadValueInt(mail, "start")
			length := module.GetPayloadValueInt(mail, "length")
			topword(start, length)
		default:
			glog.Warning("DB unhandler mail cmd:", cmd)
		}
	default:
		glog.Warning("DB unhandler mail type:", mail.Type)
	}

}
