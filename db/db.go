package db

import (
	"github.com/golang/glog"
)

func saveword(nickName, word string) error {
	glog.Info("DB save word ", nickName, word)
	sql := newstatement().insert("word").columns("nickName", "word").values(nickName, word).toString()
	if suc, err := Exec(sql); !suc {
		glog.Warning("DB saveword:", err)
		return err
	}

	return nil
}

func topword(start, len int) {
	glog.Info("DB top word ", start, len)
}
