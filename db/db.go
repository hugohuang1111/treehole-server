package db

import (
	"github.com/golang/glog"
)

func saveword(nickName, word string) {
	glog.Info("DB save word ", nickName, word)

}

func topword(start, len int) {
	glog.Info("DB top word ", start, len)
}
